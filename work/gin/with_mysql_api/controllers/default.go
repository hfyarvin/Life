package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultGet(c *gin.Context) {
	c.String(http.StatusOK, "Hello , it's find!")
}
