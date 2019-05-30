package user

import (
	"github.com/Bowser1704/go-study/gin-demo/handler"
	"github.com/Bowser1704/go-study/gin-demo/model"
	"github.com/Bowser1704/go-study/gin-demo/pkg/errno"
	"github.com/apiserver_demos/demo07/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {
	log.Infof("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	//get user id from url parameter
	userId, _ := strconv.Atoi(c.Param("id"))

	//binding thr user data bind --> 捆绑
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	//update user by the user id
	u.Id = uint64(userId)

	//validate 校验
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
