package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func SignMD5WithRSA(privateKey *rsa.PrivateKey, s string) (string, error) {
	hashed := md5.Sum([]byte(s))
	signer, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signer), nil
}

func VerifyMD5WithRSA(publicKey *rsa.PublicKey, s, sign string) error {
	inSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	hashed := md5.Sum([]byte(s))
	return rsa.VerifyPKCS1v15(publicKey, crypto.MD5, hashed[:], inSign)
}

func SignSHA256WithRSA(privateKey *rsa.PrivateKey, s string) (string, error) {
	hashed := sha256.Sum256([]byte(s))
	signer, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signer), nil
}

func VerifySHA256WithRSA(publicKey *rsa.PublicKey, s, sign string) error {
	inSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256([]byte(s))
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], inSign)
}
