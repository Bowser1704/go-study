package user

import (
	"github.com/Bowser1704/go-study/gin-demo/handler"
	"github.com/Bowser1704/go-study/gin-demo/model"
	"github.com/Bowser1704/go-study/gin-demo/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Get(c *gin.Context) {
	//c.Param("") is string string -> int by strconv.Atoi()
	Id,_ := strconv.Atoi(c.Param("id"))

	user, err := model.Getuser(uint64(Id))
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}