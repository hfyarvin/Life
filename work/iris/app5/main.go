package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type HelloWorldController struct {
	mvc.C //继承
}

func main() {
	app := iris.New()
	app.Controller("./helloworld", new(HelloWorldController))
	app.Run(iris.Addr("localhost:7001"))
}

func (c *HelloWorldController) Get() string {
	return "This is my default action..."
}

func (c *HelloWorldController) GetBy(name string) string {
	return "Hello " + name
}

func (c *HelloWorldController) GetWelcome() (string, int) {
	return "This is the GetWelcome action func...", iris.StatusOK
}

func (c *HelloWorldController) GetWelcomeBy(name string, numTimes int) {
	c.Ctx.Writef("Hello %s, NumTimes is: %d", name, numTimes)
}
