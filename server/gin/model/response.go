package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 200
)

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context, msg string, data interface{}) {
	Result(c, SUCCESS, msg, data)
}

func Fail(c *gin.Context, msg string, data interface{}) {
	Result(c, ERROR, msg, data)
}
