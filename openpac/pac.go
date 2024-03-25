package openpac

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/zmicro-team/pab/extrand"
	"golang.org/x/sync/singleflight"

	"github.com/zmicro-team/pab/cert"
)

type Config struct {
	URL               string // api url
	AppID             string // app id
	PublicCertPath    string // 银行平台公钥
	PfxPath           string // 商户私钥
	PfxPwd            string // 传 "1"
	Dn                string // openssl x509 -inform (pem|der) -subject -nameopt RFC2253 -noout -in conf/open.cer
	CnsmrUserShortNo  string // CnsmrSeqNo 用户短号, 6位
	FundSummaryAcctNo string // 交易总账户
	MrchCode          string // 商户编号
	TraderNo          string // 商户在云收款系统的编号
	// url
	CashierDeskURL    string // 收银台地址
	FrontSkipURL      string // 前端跳转url
	CallBackNoticeURL string // 后端跳转url
}

type Client struct {
	pk         string
	sdkType    string
	publicKey  *rsa.PublicKey  // 平台公钥
	privateKey *rsa.PrivateKey // 商户私钥
	Config

	accessToken atomic.Value
	group       singleflight.Group

	httpc *http.Client // default use http.DefaultClient
}

func NewClient(c Config) (*Client, error) {
	publicKey, err := cert.LoadRSAPublicKeyFromFile(c.PublicCertPath)
	if err != nil {
		return nil, err
	}
	privateKey, cc, err := cert.LoadPfxFromFile(c.PfxPath, c.PfxPwd)
	if err != nil {
		return nil, err
	}

	c.Dn = dnEscape(c.Dn)
	client := &Client{
		pk:         pkEscape(base64.StdEncoding.EncodeToString(cc.RawSubjectPublicKeyInfo)),
		sdkType:    "api",
		publicKey:  publicKey,
		privateKey: privateKey,
		Config:     c,

		httpc: http.DefaultClient,
	}
	client.accessToken.Store("")
	go client.refreshAccessToken()
	return client, nil
}

func (c *Client) GetCashierDeskURL() string    { return c.CashierDeskURL }
func (c *Client) GetFrontSkipURL() string      { return c.FrontSkipURL }
func (c *Client) GetCallBackNoticeURL() string { return c.CallBackNoticeURL }
func (c *Client) GetMrchCode() string          { return c.MrchCode }
func (c *Client) GetTraderNo() string          { return c.TraderNo }

func (c *Client) refreshAccessToken() (string, error) {
	// in flight
	tk, err, _ := c.group.Do("refresh_token", func() (any, error) {
		// 置缓存token失效,让新进来的直接进入刷新token
		c.accessToken.Store("")
		// 获取token
		tk, err := c.getAppAccessToken(context.Background())
		if err != nil {
			return "", err
		}
		// 更新缓存
		c.accessToken.Store(tk)
		return tk, nil
	})
	return tk.(string), err
}

func (c *Client) getAccessToken() (string, error) {
	if token := c.accessToken.Load().(string); token != "" {
		return token, nil
	}
	return c.refreshAccessToken()
}

func (c *Client) InvokeCloudPay(ctx context.Context, serviceID string, req, res any) error {
	mp, err := MapStructEncode(req)
	if err != nil {
		return err
	}
	mp["TraderNo"] = c.TraderNo
	rsp, err := c.invoke(ctx, serviceID, mp)
	if err != nil {
		return err
	}
	return MapStructWeakDecode(rsp, res)
}

func (c *Client) InvokeJZB(ctx context.Context, serviceID string, req, res any) error {
	mp, err := MapStructEncode(req)
	if err != nil {
		return err
	}

	mp["FundSummaryAcctNo"] = c.FundSummaryAcctNo
	rsp, err := c.invoke(ctx, serviceID, mp)
	if err != nil {
		return err
	}
	return MapStructWeakDecode(rsp, res)
}

func (c *Client) invoke(ctx context.Context, serviceID string, mp map[string]any) (map[string]any, error) {
	mp["CnsmrSeqNo"] = generateCnsmrSeqNo(c.CnsmrUserShortNo)
	mp["MrchCode"] = c.MrchCode

	for retry := 0; retry < 1; retry++ {
		result, err := c.doRequest(ctx, serviceID, mp)
		if err != nil {
			return nil, err
		}
		// try to refresh token
		// m["TxnReturnCode"] == "000000", 成功
		if _, exist := result["tokenExpiryFlag"]; exist {
			flag := result["tokenExpiryFlag"]
			delete(result, "tokenExpiryFlag")
			if flag == "true" {
				c.refreshAccessToken()
				continue
			}
		}
		return result, nil
	}
	return nil, errors.New("请求失败, 请稍后重试")
}

// getAppAccessToken 验证开发者,获取token
func (c *Client) getAppAccessToken(ctx context.Context) (string, error) {
	mp := map[string]any{
		"ApplicationID": c.AppID,
		"RandomNumber":  extrand.Number(6),
		"SDKType":       c.sdkType,
	}

	sign, err := Sign(c.privateKey, mp)
	if err != nil {
		return "", err
	}
	mp["RsaSign"] = signEscape(sign)
	mp["DN"] = c.Dn
	mp["PK"] = c.pk

	req, err := json.Marshal(mp)
	if err != nil {
		return "", err
	}
	url := c.URL + "/api/approveDev"

	log.Printf("request url: %s\n", url)
	log.Printf("request body: %s\n", string(req))
	rsp, err := c.doPost(ctx, url, req)
	if err != nil {
		return "", err
	}
	log.Printf("response: %s\n", rsp)

	mp = make(map[string]any)
	err = json.Unmarshal(rsp, &mp)
	if err != nil {
		return "", err
	}
	mp = CheckSign(c.publicKey, mp)

	tk, ok := mp["appAccessToken"].(string)
	if !ok || tk == "" {
		return "", errors.New("验证开发者执行失败, token获取失败")
	}
	return tk, nil
}

func (c *Client) doRequest(ctx context.Context, serviceID string, mp map[string]any) (map[string]any, error) {
	token, err := c.getAccessToken()
	if err != nil {
		return nil, err
	}

	mp["ApplicationID"] = c.AppID
	mp["SDKType"] = c.sdkType
	mp["RandomNumber"] = extrand.Number(6)
	mp["AppAccessToken"] = token
	mp["RequestMode"] = "json"
	mp["ValidTerm"] = "20170801"
	mp["TxnTime"] = time.Now().Format("20060102150405") + "000"

	sign, err := Sign(c.privateKey, mp)
	if err != nil {
		return nil, err
	}
	mp["RsaSign"] = sign

	url := c.URL + "/api/group/" + serviceID
	req, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}
	log.Printf("request url: %s\n", url)
	log.Printf("request body: %s\n", string(req))

	rsp, err := c.doPost(ctx, url, req)
	if err != nil {
		return nil, err
	}
	log.Printf("response: %s\n", rsp)

	mp = make(map[string]any)
	err = json.Unmarshal(rsp, &mp)
	if err != nil {
		return nil, err
	}
	mp = CheckSign(c.publicKey, mp)
	return mp, nil
}

func (c *Client) doPost(ctx context.Context, url string, data []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/xml;charset=UTF-8;")

	resp, err := c.httpc.Do(req)
	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// ------------ helper -----------

func generateCnsmrSeqNo(userShortNo string) string {
	return userShortNo + time.Now().Format("060102") + extrand.Number(10)
}

func signEscape(sign string) string {
	sign = strings.ReplaceAll(sign, "\r", "")
	sign = strings.ReplaceAll(sign, "\n", "")
	sign = strings.ReplaceAll(sign, "+", "%2B")
	sign = strings.ReplaceAll(sign, "=", "%3D")
	return sign
}

func pkEscape(pk string) string {
	return strings.ReplaceAll(pk, "+", "%2B")
}

func dnEscape(dn string) string {
	return strings.ReplaceAll(dn, "=", "%3D")
}
