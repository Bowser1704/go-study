package user

import (
	"github.com/Bowser1704/go-study/gin-demo/handler"
	"github.com/Bowser1704/go-study/gin-demo/model"
	"github.com/apiserver_demos/demo07/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
