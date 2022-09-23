package openpab

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/things-go/pab/filecipher"
	"github.com/things-go/pab/openpab/jwt"
)

/**
  平安下载的对账文件是加密的.
  <1>: 原始数据 {xxxx}.txt.enc/{xxxx}.ok.enc zip 压缩数据 并且加密的
  <2>: <1> 解密后 {xxxx}.txt.enc/{xxxx}.ok.enc 的 zip 压缩数据
  <3>: <2>解压后为 {xxxx}.txt.enc/{xxxx}.ok.enc 数据
*/

// FileDownload 下载文件, 原始文件, 即使 <1>
// filename 远端文件路径
func (c *Client) FileDownload(ctx context.Context, filename string) (io.ReadCloser, error) {
	if filename == "" {
		return nil, errors.New("文件名不能为空")
	}
	if !c.TestAccessToken() {
		c.RefreshToken(ctx)
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

// FileDownloadOriginalFile 下载文件, 原始文件, 保存到本地 即 <1>
// filename 远端文件路径
// saveFilename 本地文件路径
func (c *Client) FileDownloadOriginalFile(ctx context.Context, filename string, saveFilename string) error {
	if filename == "" || saveFilename == "" {
		return errors.New("文件名或者保存文件名不能为空")
	}
	dir := filepath.Dir(saveFilename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	r, err := c.FileDownload(ctx, filename)
	if err != nil {
		return err
	}
	defer r.Close()

	tmpFileName := saveFilename + ".tmp"
	err = WriteFileFromReader(tmpFileName, r, 0644)
	if err != nil {
		return err
	}
	return os.Rename(tmpFileName, saveFilename)
}

// FileDownloadZip 下载文件, 并解密成zip数据, 即 <2>
// filename 远端文件路径
func (c *Client) FileDownloadZip(ctx context.Context, filename, randomPassword string) ([]byte, error) {
	r, err := c.FileDownload(ctx, filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return filecipher.Decrypt2EncZipFromReader(r, randomPassword)
}

// FileDownloadZip 下载文件, 并解密成zip并保存文件, 即 <2>
// filename 远端文件路径
func (c *Client) FileDownloadZipFile(ctx context.Context, filename, randomPassword, saveFilename string) error {
	if err := os.MkdirAll(filepath.Dir(saveFilename), os.ModePerm); err != nil {
		return err
	}
	if !strings.HasSuffix(saveFilename, ".zip") {
		saveFilename += ".zip"
	}
	data, err := c.FileDownloadZip(ctx, filename, randomPassword)
	if err != nil {
		return err
	}
	return os.WriteFile(saveFilename, data, 0644)
}

// FileDownloadRaw 下载文件, 解密且解压zip, 获取原始文本数据 即使 <3>
// filename 远端文件路径
func (c *Client) FileDownloadRaw(ctx context.Context, filename, randomPassword string) ([]byte, error) {
	r, err := c.FileDownload(ctx, filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return filecipher.Decode2EncFromReader(r, filename, randomPassword)
}

// FileDownloadRaw 下载文件, 解密且解压zip获取文本数据并保存成文件 即使 <3>
// filename 远端文件路径
func (c *Client) FileDownloadRawFile(ctx context.Context, filename, randomPassword, saveFilename string) error {
	if err := os.MkdirAll(filepath.Dir(saveFilename), os.ModePerm); err != nil {
		return err
	}
	data, err := c.FileDownloadRaw(ctx, filename, randomPassword)
	if err != nil {
		return err
	}
	return os.WriteFile(saveFilename, data, 0644)
}

func WriteFileFromReader(filename string, r io.Reader, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
