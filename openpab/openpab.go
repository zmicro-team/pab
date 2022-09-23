package openpab

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/things-go/extrand"
	"github.com/tidwall/gjson"

	"github.com/things-go/pab/gm/sm2"
	"github.com/things-go/pab/openpab/jwt"
	"github.com/things-go/pab/openpab/util/idgen"
	"github.com/things-go/pab/plog"
	"github.com/things-go/pab/trace"

	"golang.org/x/sync/singleflight"
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
}

type Client struct {
	config     Config
	privateKey *sm2.PrivateKey
	publicKey  *sm2.PublicKey

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

// WithLogger 设置自定义日志
func WithLogger(l plog.Logger) Option {
	return func(c *Client) {
		c.log = l
	}
}

// WithHTTPClient 设置 http client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

// WithTraceId 设置追踪id获取, 默认框架自己生成, 可自己设置从上下文获取
func WithTraceId(f func(ctx context.Context) string) Option {
	return func(c *Client) {
		if f != nil {
			c.getTraceId = f
		}
	}
}

// WithTraceHandler 追踪回调, 表示无回调
func WithTraceHandler(f func(ctx context.Context, interId string, req, resp []byte, err error)) Option {
	return func(c *Client) {
		c.traceHandler = f
	}
}

// WithInterestTraceInterId 当追踪回调有效时, 设置感兴趣的接口id
// 默认是对所有接口都感兴趣
func WithInterestTraceInterId(ss ...string) Option {
	return func(c *Client) {
		if c.interestTraceInterId == nil {
			c.interestTraceInterId = make(map[string]struct{})
		}
		for _, v := range ss {
			c.interestTraceInterId[v] = struct{}{}
		}
	}
}

func New(config Config, opts ...Option) *Client {
	client := &Client{
		config:     config,
		httpClient: &http.Client{Timeout: time.Second * 10},
		log:        plog.NewStdLogger(os.Stderr),
		getTraceId: func(ctx context.Context) string { return trace.NextTraceId() },
	}
	client.accessToken.Store("") // set atomic value must be string
	for _, opt := range opts {
		opt(client)
	}

	curve := sm2.P256Sm2()
	client.privateKey = new(sm2.PrivateKey)
	client.privateKey.D, _ = new(big.Int).SetString(config.PrivateKey, 16)

	client.privateKey.PublicKey.Curve = curve
	client.privateKey.PublicKey.X, client.privateKey.PublicKey.Y = curve.ScalarBaseMult(client.privateKey.D.Bytes())

	client.publicKey = new(sm2.PublicKey)
	client.publicKey.Curve = curve
	client.publicKey.X, _ = new(big.Int).SetString(config.PublicKey[2:66], 16)
	client.publicKey.Y, _ = new(big.Int).SetString(config.PublicKey[66:], 16)

	go client.RefreshToken(context.Background())

	return client
}

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
		result, err := c.getCredentialsToken(ctx)
		if err != nil {
			c.log.Errorf("获取Token失败, %v", err)
			return "", err
		}
		c.setAccessToken(result.AccessToken)
		return result.AccessToken, nil
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

// Invoke 处理请求, req 只需要非公共字段, 底层已加入公共字段
func (c *Client) Invoke(ctx context.Context, interId string, req, resp any) error {
	var err error
	var respBody []byte

	serviceId, err := GetServiceId(interId)
	if err != nil {
		return err
	}
	reqBody, err := c.encodeBody(interId, req)
	if err != nil {
		return err
	}
	if c.needDoTraceHandler(interId) {
		defer func() {
			c.traceHandler(ctx, interId, reqBody, respBody, err)
		}()
	}
	respBody, err = c.invoke(ctx, serviceId, reqBody)
	if err != nil {
		return err
	}
	return json.Unmarshal(respBody, resp)
}

func (c *Client) invoke(ctx context.Context, serviceId string, req []byte) ([]byte, error) {
	var err error
	var encrypt bool // 是否加密

	reqBody := string(req)
	reqBody, err = c.encryptBody(reqBody, encrypt)
	if err != nil {
		return nil, err
	}
	if !c.TestAccessToken() {
		c.RefreshToken(ctx)
	}

	for retry := 0; retry < 2; retry++ {
		headers := c.encodeHeader(nil, reqBody, encrypt)
		result, err := c.post(ctx, c.config.BaseUrl+"/V1.0/"+serviceId, headers, reqBody)
		if err != nil {
			return nil, err
		}
		respBody, err := c.decodeBody(result, encrypt)
		if err != nil {
			return nil, err
		}

		// {"Code":"E30001","Message":"Authorization头字段缺失或令牌无效","Errors":[{"ErrorCode":"OB.Auth.Checktoken&Singed&decrypt.Failed","ErrorMessage":"Authorization头字段缺失或令牌无效"}]}
		if gjson.GetBytes(respBody, "Code").String() == "E30001" {
			c.RefreshToken(ctx)
			continue
		}
		return respBody, nil
	}
	return nil, errors.New("请求失败, 请稍后重试")
}

func (c *Client) encodeHeader(clientHeaders map[string]string, body string, encrypt bool) map[string]string {
	sig, err := jwt.SignJWS(c.config.AppId, c.privateKey, body)
	if err != nil {
		return nil
	}
	headers := map[string]string{
		"x-pab-signature":    sig,
		"x-pab-appID":        c.config.AppId,
		"x-pab-global-seqno": fmt.Sprintf("%d", idgen.Next()),
		"x-pab-timestamp":    time.Now().Format("2006-01-02 15:04:05"),
		"x-pab-version":      "GO-BOAP_1.0.0",
		"x-pab-signMethod":   "SM2",
		"x-pab-encrypt":      strconv.FormatBool(encrypt),
	}
	if tk := c.GetAccessToken(); tk != "" {
		headers["Authorization"] = tk
	}
	if encrypt {
		headers["x-pab-encryptMethod"] = "SM4"
	}
	if len(clientHeaders) > 0 {
		b, _ := json.Marshal(clientHeaders)
		headers["x-client-headers"] = string(b)
	}
	return headers
}

func (c *Client) encryptBody(b string, encrypt bool) (string, error) {
	var err error

	if !encrypt {
		return b, nil
	}
	b, err = encryptDataWithSM4(c.config.AppId, c.publicKey, b)
	if err != nil {
		return "", fmt.Errorf("openpab: 加密失败, %w", err)
	}
	return b, nil
}

func (c *Client) encodeBody(interId string, req any) ([]byte, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	reqBody := make(map[string]any)
	err = json.Unmarshal(reqBytes, &reqBody)
	if err != nil {
		return nil, err
	}

	commonBody := map[string]any{
		// 公共字段
		"CnsmrSeqNo": c.generateCnsmrSeqNo(22), // 建议: 用户短号(6位)+日期(6位)+随机编号(10位)
		"TxnCode":    interId,                  // 交易码/接口id
		// "TxnClientNo":       "", // 交易客户号, Ecif客户号, not used
		"TxnTime":  time.Now().Format("20060102150405") + "000", // 发送时间, 格式: YYYYMMDDHHmmSSNNN, 后三位固定0
		"MrchCode": c.config.MrchCode,                           // 商户号, 平台代码
		// 主体
		"FundSummaryAcctNo": c.config.FundSummaryAcctNo, // 资金汇总账号
	}
	for k, v := range commonBody {
		reqBody[k] = v
	}
	return json.Marshal(reqBody)
}

func (c *Client) decodeBody(resp *HttpResponse, encrypt bool) (body []byte, err error) {
	body = resp.Body()
	sig := ""
	headers := resp.Header()
	if h := headers["x-pab-signature"]; len(h) != 0 {
		sig = h[0]
	} else if h = headers["X-Pab-Signature"]; len(h) != 0 {
		sig = h[0]
	}
	if sig != "" {
		data := string(body)
		if resp.Size() != 0 && encrypt {
			data, err = decryptDataWithSM4(data, c.privateKey)
			if err != nil {
				return nil, err
			}
			body = []byte(data)
		}

		ok := jwt.CheckJWS(sig, c.publicKey, data)
		if !ok {
			return nil, errors.New("openpab: 签名验证失败")
		}
	}
	return body, nil
}

// generateCnsmrSeqNo 只要满足22位即可
// 建议: 用户短号(6位)+日期(6位)+随机编号(10位)
func (c *Client) generateCnsmrSeqNo(n int) string {
	return c.config.CnsmrSeqNoPrefix + time.Now().Format("060102") + extrand.Number(n-6-len(c.config.CnsmrSeqNoPrefix))
}
