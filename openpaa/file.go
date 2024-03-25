package openpaa

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/zmicro-team/pab/filecipher"
)

/**
  平安下载的对账文件是加密的.
  <1>: 原始数据 {xxxx}.txt.enc/{xxxx}.ok.enc zip 压缩数据 并且加密的
  <2>: <1> 解密后 {xxxx}.txt.enc/{xxxx}.ok.enc 的 zip 压缩数据
  <3>: <2>解压后为 {xxxx}.txt.enc/{xxxx}.ok.enc 数据
*/

// FileMsgFlag 标志
const (
	FileMsgFlagPutAuth = "101"
	FileMsgFlagGetAuth = "102"
	FileMsgFlagPut     = "201"
	FileMsgFlagGet     = "202"
	FileMsgFlagSuccess = "000000"
)

// FilePieceMerge 文件分片合并接口
type FilePieceMerge interface {
	// 写分配内容
	Write(b []byte) (int, error)
	// 完成所有分配写入, 并校验md5
	Finish(md5Value string) error
}

func (c *Client) fileSendRequest(ctx context.Context, data []byte) (io.ReadCloser, error) {
	return c.SendXmlRequest(ctx, c.config.FileDownloadUrl, data)
}

func (c *Client) FileDownload(ctx context.Context, pm FilePieceMerge, rd ReconciliationDocumentQueryItem) error {
	config := c.config

	if !c.TestAccessToken() {
		c.RefreshToken(ctx)
	}

	result, err := c.fileAuth(ctx, &FileAuthRequest{
		AppID:       config.AppId,
		Token:       c.GetAccessToken(),
		PrivateAuth: rd.DrawCode,
		Uid:         config.CnsmrSeqNoPrefix,
		Passwd:      config.CnsmrSeqNoPrefix,
		FileMsgFlag: FileMsgFlagGetAuth,
		FileName:    rd.FilePath,
	})
	if err != nil {
		return err
	}

	md5 := ""
	for index := 1; ; index++ {
		filePieceResp, err := c.filePiece(ctx, &FilePieceRequest{
			PrivateAuth: rd.DrawCode,
			Uid:         config.CnsmrSeqNoPrefix,
			FileMsgFlag: FileMsgFlagGet,
			ServerName:  result.ServerName,
			SessionID:   result.SessionID,
			FileName:    rd.FilePath,
			FileIndex:   index,
			StartPiece:  result.StartPiece,
			PieceNum:    result.PieceNum,
		})
		if err != nil {
			return err
		}
		_, err = pm.Write(filePieceResp.Content)
		if err != nil {
			return err
		}
		md5 = filePieceResp.Md5
		if filePieceResp.LastPiece {
			break
		}
	}
	return pm.Finish(md5)
}

// FileDownloadOriginalFile 下载文件, 原始文件, 保存到本地 即 <1>
// saveFilename 本地文件路径
func (c *Client) FileDownloadOriginalFile(ctx context.Context, rd ReconciliationDocumentQueryItem, saveFilename string) error {
	if saveFilename == "" {
		return errors.New("保存文件名不能为空")
	}
	dir := filepath.Dir(saveFilename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	pm, err := NewPieceMergeFile(saveFilename)
	if err != nil {
		return err
	}
	return c.FileDownload(ctx, pm, rd)
}

// FileDownloadZip 下载文件,  并解密成zip数据, 即 <2>
func (c *Client) FileDownloadZip(ctx context.Context, rd ReconciliationDocumentQueryItem) ([]byte, error) {
	pm := NewPieceMergeBuffer()
	err := c.FileDownload(ctx, pm, rd)
	if err != nil {
		return nil, err
	}
	return filecipher.Decrypt2EncZipFromReader(pm.Reader(), rd.RandomPassword)
}

// FileDownloadZipFile 下载文件,  并解密成zip并保存文件, 即 <2>
func (c *Client) FileDownloadZipFile(ctx context.Context, rd ReconciliationDocumentQueryItem, saveFilename string) error {
	if err := os.MkdirAll(filepath.Dir(saveFilename), os.ModePerm); err != nil {
		return err
	}
	if !strings.HasSuffix(saveFilename, ".zip") {
		saveFilename += ".zip"
	}
	data, err := c.FileDownloadZip(ctx, rd)
	if err != nil {
		return err
	}
	return os.WriteFile(saveFilename, data, 0644)
}

// FileDownloadRaw 下载文件, 解密且解压zip, 获取原始文本数据 即使 <3>
func (c *Client) FileDownloadRaw(ctx context.Context, rd ReconciliationDocumentQueryItem) ([]byte, error) {
	pm := NewPieceMergeBuffer()
	err := c.FileDownload(ctx, pm, rd)
	if err != nil {
		return nil, err
	}
	return filecipher.Decode2EncFromReader(pm.Reader(), rd.FileName, rd.RandomPassword)
}

// FileDownloadRawFile 下载文件, 解密且解压zip获取文本数据并保存成文件,  即 <3>
func (c *Client) FileDownloadRawFile(ctx context.Context, rd ReconciliationDocumentQueryItem, saveFilename string) error {
	if err := os.MkdirAll(filepath.Dir(saveFilename), os.ModePerm); err != nil {
		return err
	}
	data, err := c.FileDownloadRaw(ctx, rd)
	if err != nil {
		return err
	}
	return os.WriteFile(saveFilename, data, 0644)
}
