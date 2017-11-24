package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.GET("/", controllers.DefaultGet)

	router.Run(":8080")
}
