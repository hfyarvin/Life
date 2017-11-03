package route

import (
	"../controllers"
	"../controllers/file_controllers"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/", controllers.DefaultGetHandler)
	router.POST("/", controllers.DefaultPostHandler)

	//分组
	userGroup := router.Group("/user")
	{
		userGroup.GET("/info/:name", controllers.GetUserNameByRoute)
	}

	//多级分组
	userPrivateGroup := userGroup.Group("/private")
	{
		userPrivateGroup.GET("/info", controllers.GetUserInfo)
	}

	router.POST("/form/test", controllers.GetFormTest)

	//绑定
	bindGroup := router.Group("/bind")
	{
		bindGroup.POST("/form/default", controllers.BindDefaultForm)
		bindGroup.POST("/json/default", controllers.BindDefaultJson)
	}

	//文件
	fileGroup := router.Group("/file")
	{
		fileGroup.POST("/upload", file_controllers.UploadFile)
	}
}
