package openpaa

import (
	"crypto/rsa"
	"encoding/json"

	"github.com/spf13/cast"

	"github.com/things-go/pab/utils"
)

func innerNeedMap(parent string, in map[string]any, out map[string]string) {
	for k, v := range in {
		switch obj := v.(type) {
		case map[string]any:
			innerNeedMap(k, obj, out)
		case []any:
			for _, vv := range obj {
				switch u := vv.(type) {
				case map[string]any:
					innerNeedMap(k, u, out)
				case string:
					out[parent+u] = u
				}
			}
		case string:
			out[parent+k] = obj
		}
	}
}

func needSignMap(in map[string]any) map[string]string {
	out := make(map[string]string)
	innerNeedMap("", in, out)
	return out
}

func struct2Map(in any) (map[string]any, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	m := make(map[string]any)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func concatMap(mp map[string]string) string {
	return utils.ConcatSortMap(mp, "=", "&") + "&"
}

func Sign(priv *rsa.PrivateKey, mp map[string]any) (string, error) {
	need := needSignMap(mp)
	return utils.SignMD5WithRSA(priv, concatMap(need))
}

// 返回数据验签
func Verify(pub *rsa.PublicKey, mp map[string]any, sign string) error {
	need := needSignMap(mp)
	return utils.VerifyMD5WithRSA(pub, concatMap(need), sign)
}

// 验证签名
func CheckSign(pub *rsa.PublicKey, body []byte) error {
	mp := make(map[string]any)
	err := json.Unmarshal(body, &mp)
	if err != nil {
		return err
	}
	sign := cast.ToString(mp["RsaSign"])
	// 非业务错误
	errorCode := cast.ToString(mp["errorCode"])
	errorMsg := cast.ToString(mp["errorMsg"])

	delete(mp, "RsaSign")
	delete(mp, "errorCode")
	delete(mp, "errorMsg")

	// errorCode = OPEN-E-000000 验证开发者成功, 其它为 非业务错误
	// len(errorCode) == 0 业务正确, 业务错误由业务层判断
	if errorCode == "OPEN-E-000000" || len(errorCode) == 0 {
		err := Verify(pub, mp, sign)
		if err != nil {
			return NewError("", err.Error())
		}
		return nil
	}
	return NewError(errorCode, errorMsg)
}
