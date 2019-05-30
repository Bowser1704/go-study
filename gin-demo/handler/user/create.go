package user

import (
	"github.com/Bowser1704/go-study/gin-demo/handler"
	//他用了. 从而不用加handler, 但是我这里没有,所以是handler.SendResponse
	"github.com/Bowser1704/go-study/gin-demo/model"
	"github.com/Bowser1704/go-study/gin-demo/pkg/errno"
	"github.com/Bowser1704/go-study/gin-demo/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	log.Infof("user function called.", lager.Data{"X-Request-Id" : util.GetReqId(c)})
	var r CreateRequest	//request的数据结构体   创建成功后要返回数据结构CreateResponse
	//Bind是gin的方法 使用自定义结构绑定表单数据请求 不用向py里那样写在这里  只支持当前没有　form 嵌套的自定义结构  https://learnku.com/docs/gin-gonic/2018/gin-readme/3819
	if err := c.Bind(&r); err!=nil{
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//validate the data 校验数据
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
	}

	//encrypt the password  hash 一下
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
	}

	//insert data to the database 重头戏,插入数据库
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
	}

	//注意这里是r 不是u 虽然还没搞懂为什么
	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, nil, rsp)
}