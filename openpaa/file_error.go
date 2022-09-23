package openpaa

const (
	// socket 异常  序号1-10
	ErrFileSocket        = "EFS0001"
	ErrFileSocketConnect = "EFS0002"
	ErrFileSocketClose   = "EFS0003"
	ErrFileSocketTimeOut = "EFS0004"
	ErrFileHttpConnect   = "EFS0005"
	ErrFileConnClose     = "EFS0006"

	// 文件异常 序号11-40
	ErrFileFileOperate        = "EFS0011"
	ErrFileFileNotFound       = "EFS0012"
	ErrFileFileUp             = "EFS0013"
	ErrFileFileDown           = "EFS0014"
	ErrFileFileClose          = "EFS0015"
	ErrFileFileDelete         = "EFS0016"
	ErrFileFileRename         = "EFS0017"
	ErrFileFileCheck          = "EFS0018"
	ErrFileFileRead           = "EFS0019"
	ErrFileFileWrite          = "EFS0020"
	ErrFileNotFile            = "EFS0021"
	ErrFileFileOverSize       = "EFS0022"
	ErrFileFileSessionTamper  = "EFS0023"
	ErrFileFileSessionOutTime = "EFS0024"
	ErrFileFileType           = "EFS0025"

	// 转码异常 序号41-50
	ErrFileConvertFile = "EFS0041"

	// 认证信息异常 序号51-60
	ErrFileAuthUserFailed  = "EFS0051"
	ErrFileAuthTokenFailed = "EFS0052"
	ErrFileAuthTokenIsNull = "EFS0053"

	// 服务器执行异常 序号61-100
	ErrFileEncrypt           = "EFS0061"
	ErrFileDecrypt           = "EFS0062"
	ErrFileLoadConfigFile    = "EFS0063"
	ErrFileSaveConfigFile    = "EFS0064"
	ErrFileHeadXml           = "EFS0065"
	ErrFileExecCmd           = "EFS0066"
	ErrFileTransfConnect     = "EFS0067"
	ErrFileReadReqLength     = "EFS0068"
	ErrFileAuthPrivateAuth   = "EFS0069"
	ErrFileFileSessionIsNull = "EFS0070"
	ErrFileFileSessionIsOK   = "EFS0071"
)

var FileErrorCode = map[string]string{
	// socket 异常
	ErrFileSocket:        "Socket错误",
	ErrFileSocketConnect: "连接服务器异常",
	ErrFileSocketClose:   "关闭Socket流异常",
	ErrFileSocketTimeOut: "Socket通讯超时出错",
	ErrFileHttpConnect:   "连接服务器异常",
	ErrFileConnClose:     "关闭连接异常",
	// 文件异常
	ErrFileFileOperate:        "文件操作异常",
	ErrFileFileNotFound:       "文件不存在",
	ErrFileFileUp:             "文件上传错误",
	ErrFileFileDown:           "下载错误",
	ErrFileFileClose:          "文件关闭错误",
	ErrFileFileDelete:         "文件删除失败",
	ErrFileFileRename:         "文件重命名出错",
	ErrFileFileCheck:          "文件Md5校验出错",
	ErrFileFileRead:           "文件读取出错",
	ErrFileFileWrite:          "文件写入出错",
	ErrFileNotFile:            "路径非文件错误",
	ErrFileFileOverSize:       "文件内容长度超限",
	ErrFileFileSessionTamper:  "文件缓存session被篡改, 需要重新发起传输请求",
	ErrFileFileSessionOutTime: "文件缓存session失效, 需要重新发起传输请求",
	ErrFileFileSessionIsNull:  "文件缓存session不存在, 需要重新发起传输请求",
	ErrFileFileSessionIsOK:    "上传文件已经存在了, 无法续传",
	ErrFileFileType:           "请压缩文件",
	// 转码异常
	ErrFileConvertFile: "转换文件出错",
	// 认证信息异常
	ErrFileAuthUserFailed:  "用户认证失败",
	ErrFileAuthTokenFailed: "token认证失败",
	ErrFileAuthTokenIsNull: "token为空, 请先认证用户",
	// 服务器执行异常
	ErrFileEncrypt:        "信息加密错误",
	ErrFileDecrypt:        "信息解密错误",
	ErrFileLoadConfigFile: "加载配置文件出错",
	ErrFileSaveConfigFile: "保存配置文件出错",
	ErrFileHeadXml:        "XML文件格式错误",
	ErrFileExecCmd:        "执行系统命令出错",
	ErrFileTransfConnect:  "链接转发服务器异常",
	ErrFileReadReqLength:  "读取请求数据长度异常",

	ErrFileAuthPrivateAuth: "私密授权码校验失败",
}

func FileErrCodeText(code string) string {
	msg := FileErrorCode[code]
	if msg == "" {
		msg = code
	} else {
		msg = code + "-" + msg
	}
	return msg
}
