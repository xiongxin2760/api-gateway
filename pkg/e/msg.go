package e

// MsgFlags code对应报错信息
var MsgFlags = map[ErrorCode]string{
	Success:         "Ok",
	NotFound:        "资源未找到",
	Error:           "服务异常，请稍后再试",
	InvalidParams:   "请求参数错误",
	StatusForbidden: "Auth fail",
}

// GetMsg 获取code对应错误详细信息
func GetMsg(code ErrorCode) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[Error]
}
