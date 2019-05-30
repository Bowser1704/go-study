package user

import (
	"github.com/Bowser1704/go-study/gin-demo/service"
	"github.com/Bowser1704/go-study/gin-demo/handler"
	"github.com/Bowser1704/go-study/gin-demo/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ListUser(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	//r.Offest, r.Limit is string
	Offest, _ := strconv.Atoi(r.Offest)
	Limit, _	:= strconv.Atoi(r.Limit)
	infos, count, err := service.Listuser(r.Username, Offest, Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount:	count,
		UserList:	infos,
	})
}
