package controllers

import (
	"../models/user_model"
	"github.com/gin-gonic/gin"
)

func BindDefaultForm(c *gin.Context) {
	var form user_model.User
	if c.Bind(&form) == nil {
		if form.User == "tom" && form.Password == "123" {
			c.JSON(200, gin.H{"status": "form logined in"})
		} else {
			c.JSON(200, gin.H{"status": "form no login"})
		}
	}
}

func BindDefaultJson(c *gin.Context) {
	var json user_model.User
	if c.BindJSON(&json) == nil {
		if json.User == "tom" && json.Password == "123" {
			c.JSON(200, gin.H{"status": "json logined in"})
		} else {
			c.JSON(201, gin.H{"status": "json no login"})
		}
	}
}
