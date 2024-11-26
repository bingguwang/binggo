package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 gin的response的封装
*/

func Ok(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}

func OkWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  msg,
	})
}

func OkWithData(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
}

func BadRequest(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"code": http.StatusBadRequest,
		"msg":  msg,
	})
}

func Error(ctx *gin.Context, msg string, datas ...map[string]interface{}) {
	res := map[string]interface{}{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	}
	if len(datas) > 0 {
		res["data"] = datas[0]
	}
	ctx.JSON(http.StatusInternalServerError, res)
}

func Denied(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		"code": http.StatusForbidden,
		"msg":  msg,
	})
}

func OkWithXmlData(ctx *gin.Context, msg string, data interface{}) {
	type result struct {
		Code int
		Msg  string
		Data interface{}
	}
	ctx.XML(http.StatusOK, &result{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	})
}
