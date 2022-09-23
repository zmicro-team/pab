package filecipher

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

// error defined
var ErrUnPaddingOutOfRange = errors.New("unPadding out of range")

func EncryptCBCTripleDES(src []byte, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	src = PCKSPadding(src, block.BlockSize())
	iv := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	cipher.NewCBCEncrypter(block, iv).CryptBlocks(src, src)
	return src, nil
}

func DecryptCBCTripleDES(cryptText, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	iv := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	origData := make([]byte, len(cryptText))

	cipher.NewCBCDecrypter(block, iv).CryptBlocks(origData, cryptText)
	return PCKSUnPadding(origData, block.BlockSize())
}

// PCKSPadding PKCS#5和PKCS#7 填充
func PCKSPadding(origData []byte, blockSize int) []byte {
	padSize := blockSize - len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(origData, padText...)
}

// PCKSUnPadding PKCS#5和PKCS#7 解填充
func PCKSUnPadding(origData []byte, blockSize int) ([]byte, error) {
	orgLen := len(origData)
	if orgLen == 0 {
		return nil, ErrUnPaddingOutOfRange
	}
	unPadSize := int(origData[orgLen-1])
	if unPadSize > blockSize || unPadSize > orgLen {
		return nil, ErrUnPaddingOutOfRange
	}
	for _, v := range origData[orgLen-unPadSize:] {
		if v != byte(unPadSize) {
			return nil, ErrUnPaddingOutOfRange
		}
	}
	return origData[:(orgLen - unPadSize)], nil
}
