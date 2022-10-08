package openpaa

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/things-go/extrand"
	"github.com/tidwall/gjson"
	"golang.org/x/sync/singleflight"

	"github.com/things-go/pab/cert"
	"github.com/things-go/pab/plog"
	"github.com/things-go/pab/trace"
)

type Config struct {
	// 商户号, 平台代码
	MrchCode string `yaml:"mrchCode" json:"mrchCode"`
	// 资金汇总账号
	FundSummaryAcctNo string `yaml:"fundSummaryAcctNo" json:"fundSummaryAcctNo"`
	// base URL
	BaseUrl string `yaml:"baseUrl" json:"baseUrl"`
	// client id
	AppId string `yaml:"appId" json:"appId"`
	// client secret
	AppSecret string `yaml:"appSecret" json:"appSecret"`
	// 获取token 的scope, 目前为空
	Scope string `yaml:"scope" json:"scope"`
	// 商户私钥, 用于解密接收平台的数据
	// sm2
	PrivateKey string `yaml:"privateKey" json:"privateKey"`
	// 平台公钥, 用于加密发往平台的数据
	// sm2
	PublicKey string `yaml:"publicKey" json:"publicKey"`
	// 文件上传url
	FileUploadUrl string `yaml:"fileUploadUrl" json:"fileUploadUrl"`
	// 文件下载url
	FileDownloadUrl string `yaml:"fileDownloadUrl" json:"fileDownloadUrl"`
	// CnsmrSeqNo 前缀, 占2~6位, 如 DD
	CnsmrSeqNoPrefix string `yaml:"cnsmrSeqNoPrefix" json:"cnsmrSeqNoPrefix"`

	// openssl x509 -inform (pem|der) -subject -nameopt RFC2253 -noout -in conf/open.cer
	Dn string `yaml:"dn" json:"dn"`
}

type Client struct {
	config     Config
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	pk         string

	httpClient  *http.Client
	accessToken atomic.Value
	group       singleflight.Group
	// 日志接口
	log                  plog.Logger
	getTraceId           func(ctx context.Context) string
	traceHandler         func(ctx context.Context, interId string, req, resp []byte, err error)
	interestTraceInterId map[string]struct{}
}

type Option func(client *Client)

func New(config Config, opts ...Option) (*Client, error) {
	var err error

	client := &Client{
		config:               config,
		privateKey:           nil,
		publicKey:            nil,
		httpClient:           &http.Client{Timeout: time.Second * 10},
		accessToken:          atomic.Value{},
		group:                singleflight.Group{},
		log:                  plog.NewStdLogger(os.Stdout),
		getTraceId:           func(ctx context.Context) string { return trace.NextTraceId() },
		traceHandler:         nil,
		interestTraceInterId: nil,
	}
	client.accessToken.Store("") // set atomic value must be string
	for _, opt := range opts {
		opt(client)
	}
	client.publicKey, err = cert.LoadRSAPublicKeyFromFile(config.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("加载公钥(%s)失败, %v", config.PublicKey, err)
	}
	priv, certificate, err := cert.LoadPfxFromFile(config.PrivateKey, "1")
	if err != nil {
		return nil, fmt.Errorf("加载私钥(%s)失败, %v", config.PrivateKey, err)
	}
	client.privateKey = priv
	client.pk = strings.ReplaceAll(
		base64.StdEncoding.EncodeToString(certificate.RawSubjectPublicKeyInfo),
		"+",
		"%2B",
	)
	client.config.Dn = strings.ReplaceAll(client.config.Dn,
		"=",
		"%3D",
	)
	return client, nil
}
func (c *Client) GetConfig() *Config { return &c.config }

// SetLogger 设置自定义日志
func (c *Client) SetLogger(l plog.Logger) *Client {
	c.log = l
	return c
}

// SetTraceId 设置追踪id获取, 默认框架自己生成, 可自己设置从上下文获取
func (c *Client) SetTraceId(f func(ctx context.Context) string) *Client {
	if f != nil {
		c.getTraceId = f
	}
	return c
}

// SetTraceHandler 设置追踪回调, 为nil时, 表示无回调
func (c *Client) SetTraceHandler(f func(ctx context.Context, interId string, req, resp []byte, err error)) *Client {
	c.traceHandler = f
	return c
}

// SetInterestTraceInterId 当追踪回调有效时, 设置感兴趣的接口id
// 默认是对所有接口都感兴趣
func (c *Client) SetInterestTraceInterId(ss ...string) *Client {
	if c.interestTraceInterId == nil {
		c.interestTraceInterId = make(map[string]struct{})
	}
	for _, v := range ss {
		c.interestTraceInterId[v] = struct{}{}
	}
	return c
}

// GetMrchCode 商户号, 平台代码
func (c *Client) GetMrchCode() string { return c.config.MrchCode }

// RefreshToken 刷新token
func (c *Client) RefreshToken(ctx context.Context) (string, error) {
	// in flight
	tk, err, _ := c.group.Do("refresh_token", func() (any, error) {
		c.setAccessToken("") // 置token失效,让新进来的直接进入刷新token
		result, err := c.GetCredentialsToken(ctx)
		if err != nil {
			c.log.Errorf("获取Token失败, %v", err)
			return "", err
		}
		c.setAccessToken(result.AppAccessToken)
		return result.AppAccessToken, nil
	})
	if err != nil {
		return "", err
	}
	return tk.(string), nil
}

func (c *Client) GetAccessToken() string  { return c.accessToken.Load().(string) }
func (c *Client) setAccessToken(s string) { c.accessToken.Store(s) }

func (c *Client) TestAccessToken() bool { return c.GetAccessToken() != "" }

func (c *Client) needDoTraceHandler(interId string) bool {
	if c.traceHandler == nil {
		return false
	}
	if c.interestTraceInterId == nil {
		return true
	}
	_, ok := c.interestTraceInterId[interId]
	return ok
}

func (c *Client) Invoke(ctx context.Context, interId string, req, resp any) error {
	var err error
	var reqBody, respBody []byte

	serviceId, err := GetServiceId(interId)
	if err != nil {
		return err
	}

	if !c.TestAccessToken() {
		c.RefreshToken(ctx)
	}
	if c.needDoTraceHandler(interId) {
		defer func() {
			c.traceHandler(ctx, interId, reqBody, respBody, err)
		}()
	}

	for retry := 0; retry < 2; retry++ {
		reqBody, err = c.encodeBody(interId, c.GetAccessToken(), req)
		if err != nil {
			return err
		}
		result, err := c.post(ctx, c.config.BaseUrl+"/api/group/"+serviceId, nil, reqBody)
		if err != nil {
			return err
		}
		respBody = result.Body()
		err = CheckSign(c.publicKey, respBody)
		if err != nil {
			if e, ok := err.(*Error); ok && IsErrInvalidToken(e.ErrorCode) {
				c.RefreshToken(ctx)
				continue
			}
			return err
		}
		if gjson.GetBytes(respBody, "Code").String() == E30001 ||
			gjson.GetBytes(respBody, "tokenExpiryFlag").String() == "true" {
			c.RefreshToken(ctx)
			continue
		}
		err = json.Unmarshal(respBody, resp)
		if err != nil {
			return err
		}
		txnReturnCode := gjson.GetBytes(respBody, "TxnReturnCode").String()
		if txnReturnCode == "000000" {
			return nil
		}
		txnReturnMsg := gjson.GetBytes(respBody, "TxnReturnMsg").String()
		return NewError(txnReturnCode, txnReturnMsg)
	}
	return errors.New("请求失败, 请稍后重试")
}

func (c *Client) encodeBody(interId, accessToken string, req any) ([]byte, error) {
	mp, err := struct2Map(req)
	if err != nil {
		return nil, err
	}
	mp["ApplicationID"] = c.config.AppId
	mp["SDKType"] = "api"
	mp["AppAccessToken"] = accessToken
	mp["RequestMode"] = "json"
	mp["ValidTerm"] = "20170801"
	mp["TxnTime"] = time.Now().Format("20060102150405") + "000" // 发送时间, 格式: YYYYMMDDHHmmSSNNN, 后三位固定0
	mp["CnsmrSeqNo"] = c.generateCnsmrSeqNo(22)
	mp["MrchCode"] = c.config.MrchCode
	mp["FundSummaryAcctNo"] = c.config.FundSummaryAcctNo
	mp["TxnCode"] = interId

	sign, err := Sign(c.privateKey, mp)
	if err != nil {
		return nil, err
	}
	mp["RsaSign"] = sign
	return json.Marshal(mp)
}

// generateCnsmrSeqNo 只要满足22位即可
// 建议: 用户短号(6位)+日期(6位)+随机编号(10位)
func (c *Client) generateCnsmrSeqNo(n int) string {
	return c.config.CnsmrSeqNoPrefix + time.Now().Format("060102") + extrand.Number(n-6-len(c.config.CnsmrSeqNoPrefix))
}
