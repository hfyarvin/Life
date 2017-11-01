package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {
	router := gin.Default()

	var defaultGetHandler = func(c *gin.Context){
		c.String(http.StatusOK, "Hello World.")
	}
	router.GET("/", defaultGetHandler)

	router.Run(":8000")
}