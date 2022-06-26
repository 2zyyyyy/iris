package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris/web/controller"
)

func main() {
	app := iris.New()
	// 设置日志级别
	app.Logger().SetLevel("debug")
	// 注册views
	app.RegisterView(iris.HTML("./web/views", ".html"))
	// 注册控制器
	mvc.New(app.Party("/pokemon")).Handle(new(controller.PokemonController))
	mvc.New(app.Party("/trainer")).Handle(new(controller.TrainerController))
	// 启动
	_ = app.Run(iris.Addr("127.0.00.1:8080"))
}
