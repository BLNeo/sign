package e

var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorAuthCheckTokenFail:    "token鉴权失败",
	ErrorAuthCheckTokenTimeout: "token have expired",
	ErrorAuthToken:             "token生成失败",
	ErrorAuth:                  "token错误",
	ErrorAuthNotExist:          "账户不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
