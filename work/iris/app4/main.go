package main

import (
	"github.com/kataras/iris"
)

var indexPage = func(ctx iris.Context) {
	ctx.ViewData("message", "Hello World!")
	ctx.View("hello.html") //模版名
}

var userIndex = func(ctx iris.Context) {
	userId, _ := ctx.Params().GetInt64("id")
	// ctx.Writef("User Id:: %d", userId)
	ctx.ViewData("userId", userId)
	ctx.View("user_index.html") //模版名
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html")) //模版位置

	app.Get("/", indexPage)

	app.Get("/user/{id:long}", userIndex)

	app.Run(iris.Addr(":7000"))
}
