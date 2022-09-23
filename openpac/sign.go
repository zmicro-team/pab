package openpac

import (
	"crypto/rsa"

	"github.com/things-go/pab/utils"
)

func getNeedMap(in map[string]any, out map[string]string) {
	for k, v := range in {
		switch obj := v.(type) {
		case map[string]any:
			getNeedMapInner(k, obj, out)
		case []any:
			for _, o := range obj {
				switch u := o.(type) {
				case string:
					out[u] = u
				case map[string]any:
					getNeedMapInner(k, u, out)
				}
			}
		case string:
			out[k] = obj
		}
	}
}

func getNeedMapInner(parent string, in map[string]any, out map[string]string) {
	for k, v := range in {
		switch obj := v.(type) {
		case map[string]any:
			getNeedMapInner(k, obj, out)
		case []any:
			for _, o := range obj {
				switch u := o.(type) {
				case string:
					out[parent+u] = u
				case map[string]any:
					getNeedMapInner(k, u, out)
				}
			}
		case string:
			out[parent+k] = obj
		}
	}
}

func Sign(privateKey *rsa.PrivateKey, request map[string]any) (string, error) {
	need := make(map[string]string)
	getNeedMap(request, need)
	return utils.SignMD5WithRSA(privateKey, concatMap(need))
}

// 返回数据验签
func Verify(publicKey *rsa.PublicKey, m map[string]any, sign string) error {
	need := make(map[string]string)
	getNeedMap(m, need)
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
