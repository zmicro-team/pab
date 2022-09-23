package filecipher

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/**
  平安下载的对账文件是加密的.
  <1>: 原始数据 {xxxx}.txt.enc/{xxxx}.ok.enc zip 压缩数据 并且加密的
  <2>: <1> 解密后 {xxxx}.txt.enc/{xxxx}.ok.enc 的 zip 压缩数据
  <3>: <2>解压后为 {xxxx}.txt.enc/{xxxx}.ok.enc 数据
*/

// Decrypt2EncZip  从字节流 解密 <1> 数据获取到 <2>
func Decrypt2EncZip(data []byte, randomPassword string) ([]byte, error) {
	secretKey, err := base64.StdEncoding.DecodeString(randomPassword)
	if err != nil {
		return nil, err
	}
	return DecryptCBCTripleDES(data, secretKey)
}

// Decrypt2EncZipFromFile 从文件 解密 <1> 数据获取到 <2>
func Decrypt2EncZipFromFile(filename, randomPassword string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Decrypt2EncZip(b, randomPassword)
}

// Decrypt2EncZipFromReader  从io.Reader 解密 <1> 数据获取到 <2>
func Decrypt2EncZipFromReader(r io.Reader, randomPassword string) ([]byte, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return Decrypt2EncZip(b, randomPassword)
}

// Decode2EncFromFile 从文件 解密 <1> 数据获取到 <3>
func Decode2EncFromFile(filename, randomPassword string) ([]byte, error) {
	data, err := Decrypt2EncZipFromFile(filename, randomPassword)
	if err != nil {
		return nil, err
	}
	zf, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	filename = strings.TrimRight(filepath.Base(filename), ".enc")
	fs, err := zf.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	return io.ReadAll(fs)
}

// Decode2EncFromReader 从io.Reader 解密 <1> 数据获取到 <3>
// filename: {xxxx}.ok.enc/{xxxx}.txt.enc
func Decode2EncFromReader(r io.Reader, filename, randomPassword string) ([]byte, error) {
	data, err := Decrypt2EncZipFromReader(r, randomPassword)
	if err != nil {
		return nil, err
	}
	zf, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	filename = strings.TrimRight(filepath.Base(filename), ".enc")
	fs, err := zf.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	return io.ReadAll(fs)
}

// Decode2CountFromFile 从文件 解密 <1> 数据获取到 <3>, 并获得 {xxxx}.ok.enc 里的的值
// filename: {xxxx}.ok.enc
func Decode2CountFromFile(filename, randomPassword string) (int64, error) {
	data, err := Decode2EncFromFile(filename, randomPassword)
	if err != nil {
		return 0, err
	}
	r := bufio.NewReader(bytes.NewReader(data))
	v, _, err := r.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(string(v), 10, 64)
}

// Decode2CountFromReader 从 io.Reader 解密 <1> 数据获取到 <3>, 并获得 {xxxx}.ok.enc 里的的值
// filename: {xxxx}.ok.enc
func Decode2CountFromReader(r io.Reader, filename, randomPassword string) (int64, error) {
	data, err := Decode2EncFromReader(r, filename, randomPassword)
	if err != nil {
		return 0, err
	}
	rr := bufio.NewReader(bytes.NewReader(data))
	v, _, err := rr.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(string(v), 10, 64)
}
