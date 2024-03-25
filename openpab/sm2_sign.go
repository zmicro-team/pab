package openpab

import (
	"github.com/zmicro-team/pab/gm/sm2"
	"github.com/zmicro-team/pab/openpab/jwt/variant"
)

func encryptDataWithSM4(kid string, publicKey *sm2.PublicKey, plainText string) (string, error) {
	return variant.EncryptJWE(kid, publicKey, plainText)
}

func decryptDataWithSM4(jwe string, privateKey *sm2.PrivateKey) (string, error) {
	return variant.DecryptJWE(jwe, privateKey)
}

func encryptData(kid string, publicKey *sm2.PublicKey, plainText string) (string, error) {
	return "", nil
}

func decryptData(jwe string, privateKey *sm2.PrivateKey) (string, error) {
	return "", nil
}

func sign(kid string, privateKey *sm2.PrivateKey, plainText string) (string, error) {
	return "", nil
}

func sortSign(kid string, privateKey *sm2.PrivateKey, plainText string) (string, error) {
	return "", nil
}

func verifySortSign(jwt string, publicKey *sm2.PublicKey, plainText string) (string, error) {
	return "", nil
}

func verifySign(jwt string, publicKey *sm2.PublicKey, plainText string) (string, error) {
	return "", nil
}
