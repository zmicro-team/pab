package openpac

import (
	"crypto/rsa"

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
func Sign(privateKey *rsa.PrivateKey, mp map[string]any) (string, error) {
	need := needSignMap(mp)
	return utils.SignMD5WithRSA(privateKey, concatMap(need))
}

// 返回数据验签
func Verify(publicKey *rsa.PublicKey, mp map[string]any, sign string) error {
	need := needSignMap(mp)
	return utils.VerifyMD5WithRSA(publicKey, concatMap(need), sign)
}

func CheckSign(publicKey *rsa.PublicKey, m map[string]any) map[string]any {
	var sign string
	var errorCode string
	var errorMsg string

	if _, exist := m["RsaSign"]; exist {
		sign = m["RsaSign"].(string)
		delete(m, "RsaSign")
	}
	if _, exist := m["errorCode"]; exist {
		errorCode = m["errorCode"].(string)
		delete(m, "errorCode")
	}
	if _, exist := m["errorMsg"]; exist {
		errorMsg = m["errorMsg"].(string)
		delete(m, "errorMsg")
	}
	if _, exist := m["sendFlag"]; exist {
		m["TxnReturnCode"] = errorCode
		m["TxnReturnMsg"] = errorMsg
		delete(m, "sendFlag")
		return m
	}

	if len(sign) > 0 {
		if err := Verify(publicKey, m, sign); err != nil {
			m = make(map[string]any)
			m["TxnReturnCode"] = "返回-200002"
			m["TxnReturnMsg"] = "返回数据rsa验签失败"
		}
		return m
	}
	if len(errorCode) > 0 {
		m["TxnReturnCode"] = errorCode
		m["TxnReturnMsg"] = errorMsg
	} else {
		m["TxnReturnCode"] = "返回-200002"
		m["TxnReturnMsg"] = "返回数据rsa验签失败"
	}

	return m
}
