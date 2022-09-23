package utils

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GbkToUtf8(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder())
	dest, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func Utf8ToGbk(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewEncoder())
	dest, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func GBKToUtf8(src string) string {
	dest, _, _ := transform.String(simplifiedchinese.GBK.NewEncoder(), src)
	return dest
}

func Utf8ToGBK(src string) string {
	dest, _, _ := transform.String(simplifiedchinese.GBK.NewEncoder(), src)
	return dest
}
