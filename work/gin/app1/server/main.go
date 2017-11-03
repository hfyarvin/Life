package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getTest1(c *gin.Context) {
	c.String(http.StatusOK, "test1 OK")
}

func postTest1(c *gin.Context) {
	c.String(http.StatusOK, "test2 OK")
}

//解析一-----------Param------------------
//get method; params:  name/passwd
func userVerification(c *gin.Context) {
	name := c.Param("name")
	passwd := c.Param("pwd")

	c.String(http.StatusOK, "name : %s; password: %s", name, passwd)
}

func newUserVerification(c *gin.Context) {
	name := c.Param("name")
	passwd := c.Param("pwd")

	c.String(http.StatusOK, "name : %s; password: %s", name, passwd)
}

func main() {
	router := gin.Default() //注册路由

	router.GET("/test1", getTest1)
	router.POST("/test2", postTest1)
	router.GET("/user/:name/:pwd", userVerification) //http://localhost:8888/user/arvin/123
	//参数可不传
	router.GET("/new/user/:name/*pwd", newUserVerification) //http://localhost:8888/new/user/arvin/
	//Query
	router.GET("/hello", helloUser) //http://localhost:8888/hello?name=arvin
	router.POST("/postform", getMessageTest)
	router.Run(":8888")
}

//解析二----------URL中传参数-------------------------------
func helloUser(c *gin.Context) {
	name := c.Query("name")                //解析'?'后参数
	pwd := c.DefaultQuery("pwd", "123456") //有默认值
	c.String(http.StatusOK, "Hello %s(%s)", name, pwd)
}

//解析三------------PostForm--------------------------------------
func getMessageTest(c *gin.Context) {
	message := c.PostForm("message")
	extra := c.PostForm("extra")
	nick := c.DefaultPostForm("nick", "arvin")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
		"extra":   extra,
	})
}

//----------------------file------------------------
// func uploadFile(c *gin) {
// file, header, err := c.re
// }
