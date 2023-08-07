package openpaa

import (
	"context"
	"encoding/xml"
	"fmt"
)

// FilePieceRequest 文件分片请求
type FilePieceRequest struct {
	XMLName     xml.Name `xml:"FileRoot"`
	PrivateAuth string   `xml:"privateAuth"`          // 下载文件的私钥授权码
	Uid         string   `xml:"Uid"`                  // 用户id
	FileMsgFlag string   `xml:"FileMsgFlag"`          // 操作标志
	ServerName  string   `xml:"ServerName"`           // 文件服务器名称
	SessionID   string   `xml:"sessionID"`            // 会话id
	FileName    string   `xml:"FileName"`             // 文件在服务器的文件名
	FileIndex   int      `xml:"FileIndex"`            // 第几个分片
	StartPiece  int      `xml:"StartPiece,omitempty"` // 分片
	PieceNum    int      `xml:"PieceNum"`             // 分片的大小
}

// FilePieceResponse 文件分片回复
type FilePieceResponse struct {
	XMLName        xml.Name `xml:"FileRoot"`
	Uid            string   `xml:"Uid"`            // 用户id
	PrivateAuth    string   `xml:"privateAuth"`    // 下载文件的私钥授权码
	FileMsgFlag    string   `xml:"FileMsgFlag"`    // 操作标志
	SessionID      string   `xml:"sessionID"`      // 会话id
	FileName       string   `xml:"FileName"`       // 文件名
	FileSize       int      `xml:"FileSize"`       // 文件总大小
	FileIndex      int      `xml:"FileIndex"`      // 第几个分片
	StartPiece     int      `xml:"startPiece"`     // 续传开始分片
	PieceNum       int      `xml:"PieceNum"`       // 分片的大小
	LastPiece      bool     `xml:"LastPiece"`      // 是否最后一个分片
	Md5            string   `xml:"Md5"`            // md5值
	ContinueFlag   bool     `xml:"continueFlag"`   // 断点续传标识 true 续传
	TmpFileLength  int64    `xml:"TmpFileLenth"`   // 临时文件大小
	LastUpdateTime int64    `xml:"LastUpdateTime"` // 最后一次访问服务器的时间
	EbcdicFlag     bool     `xml:"EbcdicFlag"`     // 编码标志
	AuthTokenFlag  bool     `xml:"authTokenFlag"`  // token认证标志
	ScrtFlag       bool     `xml:"ScrtFlag"`       // 加密的标志
	AuthFlag       bool     `xml:"AuthFlag"`       // 认证标志
	ContLen        int      `xml:"ContLen"`        // 当前传输内容的大小
	Content        []byte   `xml:"-"`              // 文件内容
}

// filePiece 文件分片请求
func (c *Client) filePiece(ctx context.Context, req *FilePieceRequest) (*FilePieceResponse, error) {
	reqBody, err := EncodeHeaderData(req)
	if err != nil {
		return nil, err
	}
	body, err := c.fileSendRequest(ctx, reqBody)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	resp := &FilePieceResponse{}
	err = DecodeHeader(body, resp)
	if err != nil {
		return nil, err
	}
	if resp.FileMsgFlag != FileMsgFlagSuccess {
		return nil, fmt.Errorf("[%s] %d piece download failure, cause by %s", resp.FileName, resp.FileIndex, FileErrCodeText(resp.FileMsgFlag))
	}
	resp.Content, err = DecodeContent(body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
