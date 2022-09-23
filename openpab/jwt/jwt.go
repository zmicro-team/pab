package jwt

import (
	"encoding/asn1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"strings"

	"github.com/things-go/pab/gm/sm2"
)

type sm2Signature struct {
	R, S *big.Int
}

func SignJWS(kid string, privateKey *sm2.PrivateKey, plainText string) (string, error) {
	header := make(map[string]any)
	header["alg"] = "SM2"
	header["typ"] = "JWT"
	header["kid"] = kid

	b, _ := json.Marshal(header)
	base64UrlEncodeHeader := base64.RawURLEncoding.EncodeToString(b)

	jwt := base64UrlEncodeHeader + "." + "e30"

	s := sm2Signature{}

	// uid 固定 1234567812345678
	if plainText != "" {
		s.R, s.S, _ = sm2.Sm2Sign(privateKey, []byte(base64UrlEncodeHeader+"."+plainText), []byte("1234567812345678"), nil)
	} else {
		s.R, s.S, _ = sm2.Sm2Sign(privateKey, []byte(jwt), []byte("1234567812345678"), nil)
	}

	b, _ = asn1.Marshal(s)
	h := hex.EncodeToString(b)
	base64UrlSignature := base64.RawURLEncoding.EncodeToString([]byte(h))

	jwt = jwt + "." + base64UrlSignature

	return jwt, nil
}

func CheckJWS(jwt string, publicKey *sm2.PublicKey, plainText string) bool {
	parts := strings.Split(jwt, ".")

	if len(parts) != 3 {
		log.Println("malformed jwt")
		return false
	}

	if parts[0] == "" || parts[1] == "" || parts[2] == "" {
		log.Println("malformed jwt")
		return false
	}

	var M string
	if plainText != "" {
		M = parts[0] + "." + plainText
	} else {
		M = parts[0] + "." + parts[1]
	}

	b, _ := base64.RawURLEncoding.DecodeString(parts[2])

	b, _ = hex.DecodeString(string(b))
	s := sm2Signature{}
	asn1.Unmarshal(b, &s)

	ok := sm2.Sm2Verify(publicKey, []byte(M), []byte("1234567812345678"), s.R, s.S)

	return ok
}
