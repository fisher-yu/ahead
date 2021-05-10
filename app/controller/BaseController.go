package controller

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	SuccessCode   = 0
	ServerErrCode = 500
	ParamsErrCode = 400
	NotFoundErrCode = 404
)

var (
	resp Resp
)

// 获取code对应的信息
func (bc BaseController) GetCodeMsg(code int) string {
	msg := make(map[int]string, 0)
	msg[SuccessCode] = "success"
	msg[ServerErrCode] = "服务器错误"
	msg[ParamsErrCode] = "参数错误"
	msg[NotFoundErrCode] = "未找到"

	if msg[code] != "" {
		return msg[code]
	}
	return ""
}

// 成功返回
func (bc BaseController) Success(ctx *gin.Context, data interface{}) {
	resp.Code = SuccessCode
	resp.Message = bc.GetCodeMsg(SuccessCode)
	resp.Data = data
	ctx.JSON(200, resp)
}

// 失败返回
func (bc BaseController) Error(ctx *gin.Context, code int, msg string) {
	resp.Data = nil
	resp.Code = code

	if msg != "" {
		resp.Message = msg
	} else {
		resp.Message = bc.GetCodeMsg(code)
	}

	ctx.JSON(200, resp)
}
