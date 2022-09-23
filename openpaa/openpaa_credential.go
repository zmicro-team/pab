package openpaa

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/things-go/extrand"
)

type CredentialResponse struct {
	RsaSign        string `json:"RsaSign"`
	AppAccessToken string `json:"appAccessToken"`
	ErrorCode      string `json:"errorCode"`
	ErrorMsg       string `json:"errorMsg"`
}

func (c *Client) GetCredentialsToken(ctx context.Context) (*CredentialResponse, error) {
	mp := map[string]any{
		"ApplicationID": c.config.AppId,
		"RandomNumber":  extrand.Number(6),
		"SDKType":       "api",
	}

	sign, err := Sign(c.privateKey, mp)
	if err != nil {
		return nil, err
	}
	sign = strings.ReplaceAll(sign, "\r", "")
	sign = strings.ReplaceAll(sign, "\n", "")
	sign = strings.ReplaceAll(sign, "+", "%2B")
	sign = strings.ReplaceAll(sign, "=", "%3D")

	mp["DN"] = c.config.Dn
	mp["PK"] = c.pk
	mp["RsaSign"] = sign

	b, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}
	result, err := c.post(ctx, c.config.BaseUrl+"/api/approveDev", nil, b)
	if err != nil {
		return nil, err
	}

	err = CheckSign(c.publicKey, result.Body())
	if err != nil {
		return nil, err
	}
	resp := &CredentialResponse{}
	err = json.Unmarshal(result.Body(), resp)
	if err != nil {
		return nil, err
	}
	if resp.ErrorCode == "OPEN-E-000000" {
		return resp, nil
	}
	return nil, errors.New(resp.ErrorMsg)
}
