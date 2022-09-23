package openpaa

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
)

var _ FilePieceMerge = (*PieceMergeFile)(nil)

// PieceMergeFile 文件分片下载至文件
type PieceMergeFile struct {
	filename string    // 文件名
	tmpFile  *os.File  // 临时文件
	md5Hash  hash.Hash // md5计算
}

// NewPieceMergeFile 创建文件分片下载并至文件
// filename: 要保存的文件
func NewPieceMergeFile(filename string) (*PieceMergeFile, error) {
	tmpFile, err := os.OpenFile(filename+".tmp", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	return &PieceMergeFile{
		filename: filename,
		tmpFile:  tmpFile,
		md5Hash:  md5.New(),
	}, nil
}

func (e *PieceMergeFile) Write(b []byte) (int, error) {
	e.md5Hash.Write(b)
	return e.tmpFile.Write(b)
}

func (e *PieceMergeFile) Finish(md5Value string) error {
	err := e.tmpFile.Close()
	if err != nil {
		return err
	}
	computeMd5 := hex.EncodeToString(e.md5Hash.Sum(nil))
	if computeMd5 != md5Value {
		return fmt.Errorf("文件md5校验失败, local: %s, remote: %s", computeMd5, md5Value)
	}
	return os.Rename(e.filename+".tmp", e.filename)
}
