package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Controller("/helloworld", new(HelloWorldController))
	app.Run(iris.Addr("localhost:6666"))
}

type HelloWorldController struct {
	mvc.C
}

func (c *HelloWorldController) Get() string {
	return "this is my default action..."
}

func (c *HelloWorldController) GetBy(name string) string {
	return "Hello " + name
}

func (c *HelloWorldController) GetWelcome() (string, int) {
	return "This is the GetWelcome action func ...", iris.StatusOK
}

func (c *HelloWorldController) GetWelcomeBy(name string, numTimes int) {
	c.Ctx.Writef("Hello %s, NumTimes is: %d", name, numTimes)
}

func (c *HelloWorldController) GetUser() (string, int) {
	return "This is a user", iris.StatusOK
}

func (c *HelloWorldController) GetUserBy(name string) string {
	return "The user name is " + name
}
