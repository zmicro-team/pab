package openpab

import (
	"context"
	"encoding/json"
	"fmt"
	"unsafe"
)

type CredentialRequest struct {
	GrantType    string `json:"grantType"`
	Scope        string `json:"scope"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
type CredentialResult struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type CredentialContent struct {
	Result CredentialResult `json:"result"`
}

type CredentialResponse struct {
	Code         string
	Message      string
	Errors       Errors
	ResponseCode int
	Content      CredentialContent `json:"content"`
	Status       bool              `json:"status"`
}

func (c *Client) getCredentialsToken(ctx context.Context) (*CredentialResult, error) {
	b, err := json.Marshal(CredentialRequest{
		GrantType:    "client_credentials",
		Scope:        c.config.Scope,
		ClientId:     c.config.AppId,
		ClientSecret: c.config.AppSecret,
	})
	if err != nil {
		return nil, err
	}
	body := *(*string)(unsafe.Pointer(&b))
	result, err := c.post(ctx, c.config.BaseUrl+"/as/token.oauth2", c.encodeHeader(nil, body, false), body)
	if err != nil {
		return nil, err
	}
	resp := CredentialResponse{}
	err = json.Unmarshal(result.Body(), &resp)
	if resp.Code != "000000" {
		return nil, fmt.Errorf("%s, %s", resp.Code, resp.Message)
	}
	return &resp.Content.Result, nil
}
