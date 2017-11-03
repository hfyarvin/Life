package controllers

import (
	"github.com/gin-gonic/gin"
)

func DefaultGetHandler(c *gin.Context) {
	c.String(200, "Hello World!")
}

func DefaultPostHandler(c *gin.Context) {
	userName := c.PostForm("name")

	obj := gin.H{
		"say": "Hello" + userName,
	}

	c.JSON(200, obj)
}

func GetFormTest(c *gin.Context) {
	// title := c.PostForm("title")
	// count := c.DefaultPostForm("count", "cnt")

	// obj := gin.H{
	// 	"title": title,
	// 	"cnt":   count,
	// }
	obj := c.Request.PostForm
	c.JSON(200, obj)
}
