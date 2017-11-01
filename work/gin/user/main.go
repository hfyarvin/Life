package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "Hello World.")
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name") //取值
		c.String(http.StatusOK, "Hello %s !", name)
	})

	router.GET("/user/:name/*action",func(c *gin.Context){ 
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		fmt.Println(name, action,message)
		c.String(http.StatusOK, message)
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Arvin")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s.%s",firstname, lastname)
	})
	router.Run(":8000")
}