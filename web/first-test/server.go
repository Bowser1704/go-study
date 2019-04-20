package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	
	//html
	app.Handle("GET", "/hello", func(ctx iris.Context){
		ctx.HTML("<h1>Hello</h1>")
	})
	//string
	app.Get("/Ping", func(ctx iris.Context) {
		ctx.WriteString("Ping")
	})
	//json
	app.Get("/hi", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"msg":"Hello bb"})
	})
	app.Run(iris.Addr(":8080"))

}
