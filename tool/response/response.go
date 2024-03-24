package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sign/tool/log"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Output setting gin.JSON
func Output(ctx *gin.Context, httpCode int, response *Response) {
	ctx.JSON(httpCode, response)
}

func Success(ctx *gin.Context, data interface{}) {
	response := &Response{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: data,
	}
	Output(ctx, SUCCESS, response)
}

// NewErrMessage  创建自定义消息的错误类型
func NewErrMessage(code int, format string, v ...interface{}) *Response {
	errResponse := &Response{
		Code: code,
		Msg:  fmt.Sprintf(format, v...),
	}
	if format == "" {
		errResponse.Msg = GetMsg(code)
	}
	return errResponse
}

// BadRequest 传入参数的错误
func BadRequest(ctx *gin.Context, format string, v ...interface{}) {
	Output(ctx, InvalidParams, NewErrMessage(InvalidParams, format, v...))
}

// InternalError 系统内部错误
func InternalError(ctx *gin.Context, format string, v ...interface{}) {
	// log记录
	log.Logger.Error(format)
	Output(ctx, ERROR, NewErrMessage(ERROR, format, v...))
}

// UnauthorizedError 未登录错误
func UnauthorizedError(ctx *gin.Context, format string, v ...interface{}) {
	Output(ctx, Unauthorized, NewErrMessage(Unauthorized, format, v...))
}
