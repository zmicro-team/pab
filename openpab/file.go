package openpab

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/things-go/pab/openpab/jwt"
)

// error defined
var ErrUnPaddingOutOfRange = errors.New("unPadding out of range")

// FileDownload 下载文件, 原始文件
// filename 远端文件路径
func (c *Client) FileDownload(ctx context.Context, filename string) (io.ReadCloser, error) {
	if filename == "" {
		return nil, errors.New("文件名不能为空")
	}

	// NOTE: 签名一定要序列化成以下格式
	// {"fileNo":"/ejzb/5609/YE2021082556091.txt.enc","appId":"9c807ecb-5785-4c8f-9e90-d54ccead026a"}
	tmpSign, err := jwt.SignJWS(c.config.AppId, c.privateKey, fmt.Sprintf(`{"fileNo":"%s","appId":"%s"}`, filename, c.config.AppId))
	if err != nil {
		return nil, err
	}
	return c.getFile(ctx, c.config.FileDownloadUrl, map[string]string{
		"appId":  c.config.AppId,
		"fileNo": filename,
		"sign":   tmpSign,
	})
}

// FileDownload2Local 下载文件, 原始文件, 保存到本地
// filename 远端文件路径
// saveFilename 本地文件路径
func (c *Client) FileDownload2Local(ctx context.Context, filename string, saveFilename string) error {
	if filename == "" || saveFilename == "" {
		return errors.New("文件名或者文件编号不能为空")
	}

	r, err := c.FileDownload(ctx, filename)
	if err != nil {
		return err
	}
	defer r.Close()

	dir := filepath.Dir(saveFilename)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	tmpFileName := saveFilename + ".tmp"
	err = WriteFileFromReader(tmpFileName, r, 0644)
	if err != nil {
		return err
	}
	return os.Rename(tmpFileName, saveFilename)
}

// FileDownload 下载文件, 并解密成zip数据
// filename 远端文件路径
func (c *Client) FileDownloadZip(ctx context.Context, filename, randomPassword string) ([]byte, error) {
	r, err := c.FileDownload(ctx, filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return DecryptFromReader(r, randomPassword)
}

// FileDownload 下载文件, 并解密且解压zip, 获取原始文本数据
// filename 远端文件路径
func (c *Client) FileDownloadOriginal(ctx context.Context, filename, randomPassword string) ([]byte, error) {
	r, err := c.FileDownload(ctx, filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return DecodeDataFromReader(r, filename, randomPassword)
}

func WriteFileFromReader(name string, r io.Reader, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

// DecodeCountFromFile 从xxx.ok.enc文件中获取账单文件数量
func DecodeCountFromFile(name, randomPassword string) (int64, error) {
	data, err := DecryptFromFile(name, randomPassword)
	if err != nil {
		return 0, err
	}
	zf, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return 0, err
	}
	filename := strings.TrimRight(filepath.Base(name), ".enc")
	fs, err := zf.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fs.Close()
	r := bufio.NewReader(fs)
	v, _, err := r.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(string(v), 10, 64)
}

// DecodeCountFromReader 从xxx.ok.enc数据流中获取账单文件数量
func DecodeCountFromReader(r io.Reader, filename, randomPassword string) (int64, error) {
	data, err := DecryptFromReader(r, randomPassword)
	if err != nil {
		return 0, err
	}
	zf, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return 0, err
	}

	filename = strings.TrimRight(filepath.Base(filename), ".enc")
	fs, err := zf.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fs.Close()
	rr := bufio.NewReader(fs)
	v, _, err := rr.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(string(v), 10, 64)
}

// DecodeDataFromFile 从xx.txt.enc文件中解密,并zip解压
// NOTE: 此时为数据流
func DecodeDataFromFile(filename, randomPassword string) ([]byte, error) {
	data, err := DecryptFromFile(filename, randomPassword)
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

// DecodeDataFromReader 从xx.txt.enc数据流中解密,并zip解压
// NOTE: 此时为数据流
func DecodeDataFromReader(r io.Reader, filename, randomPassword string) ([]byte, error) {
	data, err := DecryptFromReader(r, randomPassword)
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

// DecryptFromFile 从xx.txt.enc或xx.ok.enc文件中解密
// NOTE: 此时为zip数据流
func DecryptFromFile(filename, randomPassword string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return DecryptEncData(b, randomPassword)
}

// DecryptFromReader 从xx.txt.enc或xx.ok.enc数据流中解密
// NOTE: 此时为zip数据流
func DecryptFromReader(r io.Reader, randomPassword string) ([]byte, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return DecryptEncData(b, randomPassword)
}

// DecryptEncData 解密数据, zip数据流
func DecryptEncData(data []byte, randomPassword string) ([]byte, error) {
	secretKey, err := base64.StdEncoding.DecodeString(randomPassword)
	if err != nil {
		return nil, err
	}
	return DecryptCBCTripleDES(data, secretKey)
}

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
