package openpaa

import (
	"context"
	"encoding/xml"
)

// FilePieceRequest 文件分片请求
type FilePieceRequest struct {
	XMLName     xml.Name `xml:"FileRoot"`
	AppID       string   `xml:"appID"`       // 门户申请的appId
	Token       string   `xml:"token"`       // token
	PrivateAuth string   `xml:"privateAuth"` // 下载文件的私钥授权码
	Uid         string   `xml:"Uid"`         // 用户id
	Passwd      string   `xml:"Passwd"`      // 用户密码
	FileMsgFlag string   `xml:"FileMsgFlag"` // 操作标志
	ServerName  string   `xml:"ServerName"`  // 文件服务器名称
	SessionID   string   `xml:"sessionID"`   // 会话id
	FileName    string   `xml:"FileName"`    // 文件在服务器的文件名
	FileIndex   int      `xml:"FileIndex"`   // 第几个分片
	StartPiece  int      `xml:"StartPiece"`  // 分片
	PieceNum    int      `xml:"PieceNum"`    // 分片的大小
}

// FilePieceResponse 文件分片回复
type FilePieceResponse struct {
	XMLName     xml.Name `xml:"FileRoot"`
	PrivateAuth string   `xml:"privateAuth"` // 下载文件的私钥授权码
	FileMsgFlag string   `xml:"FileMsgFlag"` // 操作标志
	SessionID   string   `xml:"sessionID"`   // 会话id
	FileName    string   `xml:"FileName"`    // 文件名
	FileSize    int      `xml:"FileSize"`    // 文件总大小
	FileIndex   int      `xml:"FileIndex"`   // 第几个分片
	StartPiece  int      `xml:"StartPiece"`  // 续传开始分片
	PieceNum    int      `xml:"PieceNum"`    // 分片的大小
	LastPiece   bool     `xml:"LastPiece"`   // 是否最后一个分片
	Md5         string   `xml:"Md5"`         // md5值
	Content     []byte   `xml:"-"`           // 文件内容
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
	resp.Content, err = DecodeContent(body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
