package main

import (
	// "./controllers"
	"./route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	route.Init(router)
	router.Run(":8081")
}
