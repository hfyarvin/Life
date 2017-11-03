package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World.")
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name") //解析
		c.String(http.StatusOK, "Hello %s !", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		fmt.Println(name, action, message)
		c.String(http.StatusOK, message)
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Arvin") //参数,有默认值
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s.%s", firstname, lastname)
	})

	router.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})

	// 'json'--'w-www-form-urlencoded'--'xml'-'form-data'
	router.POST("/form_post", func(c *gin.Context) {
		messgae := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},

			"message": messgae,
			"nick":    nick,
		})
	})
	/*上传多个文件
	curl -X POST http://127.0.0.1:8000/upload
	-F "upload=@e:/work/golang_personal/Life/work/gin/user/README.md"
	-H "Content-Type: multipart/form-data"
	*/
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println("the name", name)

		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}

		filename := header.Filename
		fmt.Println("the file", file, err, filename)

		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusCreated, "upload successful")
	})

	router.Run(":8000")
}
