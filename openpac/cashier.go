package openpac

import (
	"errors"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"

	"github.com/things-go/pab/utils"
)

const tpl = `
<html>
	<head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/></head>
	<body>
		<form id = "form" action="{{.ReqURL}}" method="post">
{{- range $k, $v := .Params}}
			<input type="hidden" name="{{ $k }}" id="{{ $k }}" value="{{ $v }}"/>
{{- end}}
		</form>
	</body>
	<script type="text/javascript">document.all.form.submit();</script>
</html>
`

type Tmpl struct {
	ReqURL string
	Params map[string]string
}

var htmlTpl, _ = template.New("default").Parse(tpl)
var notifySuccess = []byte("notify_success")
var decoder = schema.NewDecoder()

func init() {
	decoder.IgnoreUnknownKeys(true)
}

// CashierDesk 云收款 html body
func (c *Client) CashierDesk(req Cashier) (string, error) {
	mp, err := struct2Map(req)
	if err != nil {
		return "", err
	}
	mp["TraderNo"] = c.TraderNo
	sign, err := utils.SignSHA256WithRSA(c.privateKey, concatMap(mp))
	if err != nil {
		return "", err
	}
	mp["Signature"] = sign

	bs := &strings.Builder{}
	err = htmlTpl.Execute(bs, &Tmpl{ReqURL: c.CashierDeskURL, Params: mp})
	if err != nil {
		return "", err
	}
	return bs.String(), nil
}

// ParsePayNotify 解析支付前后台通知
func (c *Client) ParsePayNotify(resp url.Values) (*PayNotify, error) {
	var result PayNotify

	err := decoder.Decode(&result, resp)
	if err != nil {
		return nil, err
	}
	mp, err := struct2Map(result)
	if err != nil {
		return nil, err
	}

	if err = c.VerifyNotify(mp); err != nil {
		return nil, err
	}
	return &result, nil
}

// VerifyNotify 通知验签
func (c *Client) VerifyNotify(mp map[string]string) error {
	sign := mp["Signature"]
	if sign == "" {
		return errors.New("签名校验失败")
	}
	delete(mp, "Signature")
	return utils.VerifySHA256WithRSA(c.publicKey, concatMap(mp), sign)
}

// AckNotify 应答后台通知
func (*Client) AckNotify(w http.ResponseWriter) {
	_ = AckNotify(w)
}

// AckNotify 应答后台通知
func AckNotify(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	_, err := w.Write(notifySuccess)
	return err
}
