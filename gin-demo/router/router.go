package router

import (
	"github.com/Bowser1704/go-study/gin-demo/handler/user"
	"net/http"

	"github.com/Bowser1704/go-study/gin-demo/handler/sd"
	"github.com/Bowser1704/go-study/gin-demo/router/middleware"

	"github.com/gin-gonic/gin"
)

//路由管理机制
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "incorrect api")
	})

	// user handler
	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
		u.GET("/:id", user.Get)
		u.PUT("/:id", user.Update)
		u.DELETE("/:id", user.Delete)
		u.GET("", user.ListUser)
	}
	// health chack handlers
	// svcd.GET("",f) 先传入一个url,再来一个函数,一一映射.
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
