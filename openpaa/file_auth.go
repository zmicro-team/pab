package openpaa

import (
	"context"
	"encoding/xml"
	"errors"
)

// FileAuthRequest 文件授权请求
type FileAuthRequest struct {
	XMLName      xml.Name `xml:"FileRoot"`
	AppID        string   `xml:"appID"`       // 门户申请的appId
	Token        string   `xml:"token"`       // token
	PrivateAuth  string   `xml:"privateAuth"` // 下载文件的私钥授权码
	Uid          string   `xml:"Uid"`         // 用户id
	Passwd       string   `xml:"Passwd"`      // 用户密码
	FileMsgFlag  string   `xml:"FileMsgFlag"` // 操作标志
	FileName     string   `xml:"FileName"`    // 文件在服务器的文件名
	FileIndex    int      `xml:"FileIndex"`
	FileSize     int      `xml:"FileSize"`
	TmpFileLenth int      `xml:"TmpFileLenth"`
	StartPiece   int      `xml:"StartPiece"`
	PieceNum     int      `xml:"PieceNum"`
}

// FileAuthResponse 文件授权回复
type FileAuthResponse struct {
	XMLName        xml.Name `xml:"FileRoot"`
	AppID          string   `xml:"appID"`         // 门户申请的appId
	Token          string   `xml:"token"`         // token
	PrivateAuth    string   `xml:"privateAuth"`   // 下载文件的私钥授权码
	Uid            string   `xml:"Uid"`           // 用户id
	Passwd         string   `xml:"Passwd"`        // 用户密码
	FileMsgFlag    string   `xml:"FileMsgFlag"`   // 操作标志
	ServerName     string   `xml:"ServerName"`    // 文件服务器名称
	SessionID      string   `xml:"sessionID"`     // 会话id
	AuthFlag       bool     `xml:"AuthFlag"`      // 文件认证标志
	AuthTokenFlag  bool     `xml:"authTokenFlag"` // token认证标志
	FileIndex      int      `xml:"FileIndex"`     // 文件索引
	FileName       string   `xml:"FileName"`      // 文件在服务器的文件名
	FileSize       int      `xml:"FileSize"`      // 文件大小
	PieceNum       int      `xml:"PieceNum"`      // 分片的大小
	StartPiece     int      `xml:"startPiece"`    // 续传开始分片
	LastPiece      bool     `xml:"LastPiece"`
	LastUpdateTime int64    `xml:"LastUpdateTime"`
	ContLen        int      `xml:"ContLen"`
	TmpFileLenth   int      `xml:"TmpFileLenth"`
	ScrtFlag       bool     `xml:"ScrtFlag"`
	EbcdicFlag     bool     `xml:"EbcdicFlag"`
	ContinueFlag   bool     `xml:"continueFlag"`
}

// fileAuth 文件授权请求
func (c *Client) fileAuth(ctx context.Context, req *FileAuthRequest) (*FileAuthResponse, error) {
	reqBody, err := EncodeHeaderData(req)
	if err != nil {
		return nil, err
	}
	r, err := c.fileSendRequest(ctx, reqBody)
	if err != nil {
		return nil, err
	}
	result := &FileAuthResponse{}
	err = DecodeHeader(r, result)
	if err != nil {
		return nil, err
	}
	if !result.AuthFlag {
		return nil, errors.New(FileErrCodeText(ErrFileAuthUserFailed))
	}
	if !result.AuthTokenFlag {
		c.RefreshToken(ctx) // TODO: 刷新token的时机?
		return nil, errors.New(FileErrCodeText(ErrFileAuthTokenFailed))
	}
	return result, nil
}
