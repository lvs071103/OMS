package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeRecordExists
	CodeUserNotExists
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeInvalidVerifyCode
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeRecordExists:    "用户已存在",
	CodeUserNotExists:   "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",
	CodeNeedLogin:       "需要登陆",
	CodeInvalidToken:    "无效的token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

const (
	StatusInvalidParam = 1001 + iota
	StatusRecordExists
	StatusRecordNotExists
)

// OmsStatusText - 返回状态码对应的文本信息
func OtherStatusText(code int) string {
	switch code {
	case StatusInvalidParam:
		return "请求参数错误"
	case StatusRecordExists:
		return "请求记录已存在"
	case StatusRecordNotExists:
		return "用户不存在"
	default:
		return "未知错误"
	}
}
