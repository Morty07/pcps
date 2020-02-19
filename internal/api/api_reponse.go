package http_api

import (
	"pcps/pcpsd/common"

	"github.com/labstack/echo"
)

type EchoContext struct {
	C echo.Context
}

type Reponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (e *EchoContext) Response(httpCode, errCode int, data interface{}) {
	e.C.JSON(httpCode, Reponse{
		Code: errCode,
		Msg:  common.GetMsg(errCode),
		Data: data,
	})
	return
}
