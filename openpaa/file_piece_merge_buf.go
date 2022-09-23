package openpaa

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
)

var _ FilePieceMerge = (*PieceMergeBuffer)(nil)

// PieceMergeBuffer 文件分片下载buffer
type PieceMergeBuffer struct {
	tmpBuf  *bytes.Buffer
	md5Hash hash.Hash // md5计算
}

// NewPieceMergeBuffer 创建文件分配下载buffer
func NewPieceMergeBuffer() *PieceMergeBuffer {
	return &PieceMergeBuffer{
		tmpBuf:  &bytes.Buffer{},
		md5Hash: md5.New(),
	}
}

func (e *PieceMergeBuffer) Write(b []byte) (int, error) {
	e.md5Hash.Write(b)
	return e.tmpBuf.Write(b)
}

func (e *PieceMergeBuffer) Finish(md5Value string) error {
	computeMd5 := hex.EncodeToString(e.md5Hash.Sum(nil))
	if computeMd5 != md5Value {
		return fmt.Errorf("文件md5校验失败, local: %s, remote: %s", computeMd5, md5Value)
	}
	return nil
}

func (e *PieceMergeBuffer) Reader() io.Reader {
	return e.tmpBuf
}
