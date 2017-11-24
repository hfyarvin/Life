package file_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	_, header, err := c.Request.FormFile("upload")
	fmt.Println("thsii ", err)
	filename := header.Filename

	obj := gin.H{
		"filename": filename,
	}

	c.JSON(200, obj)
}
