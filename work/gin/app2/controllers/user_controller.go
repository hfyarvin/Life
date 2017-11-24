package controllers

import (
	// "fmt"
	"../models/user_model"
	"github.com/gin-gonic/gin"
)

func GetUserNameByRoute(c *gin.Context) {
	userName := c.Param("name")

	obj := gin.H{
		"username": userName,
	}
	c.JSON(200, obj)
}

func GetUserInfo(c *gin.Context) {
	// age := c.DefaultQuery("age", "18")
	// desc := ""
	// if age != "18" {
	// 	age = "18"
	// 	desc = "always 18"
	// }

	// obj := gin.H{
	// 	"user": "arvin",
	// 	"age":  age,
	// 	"desc": desc,
	// }

	obj := c.Request.URL.Query()

	c.JSON(200, obj)
}
