package handler

import (
	"github.com/Bowser1704/go-study/gin-demo/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code	int			`json:"code"`
	Message	string		`json:"message"`
	Data	interface{}	`json:"data"`	//interface{} 任意类型,感觉interface就是用来降低强类型的做法
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	//always return http.StatusOK 是神马鬼
	c.JSON(http.StatusOK, Response{
		Code:		code,
		Message:	message,
		Data:		data,
	})
}
