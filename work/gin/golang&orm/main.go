package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"strconv"
)

func main() {
	r := gin.Default()
	router(r)
	//连接数据库
	// var err errosr
	engine, err := xorm.NewEngine("mysql", "root:12345678@tcp(127.0.0.1:33.6)/yiibaidb?charset=utf8")
	if err != nil {
		fmt.Println("engine err:", err)
	}
	engine.Ping()
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	r.GET("/item/:id", func(c *gin.Context) {
		var item Item
		id := c.Param("id")
		itemId, _ := strconv.ParseInt(id, 10, 64)
		has, err := engine.Id(itemId).Get(&item)
		if err != nil {
			fmt.Println("find err:", err)
		}
		fmt.Println(has)
		fmt.Println(item)
		c.JSON(200, item)
	})
	r.Run(":9999")
}

func DefaultGet(c *gin.Context) {
	fmt.Println("i am in the function")
	obj := gin.H{
		"message": "pong",
	}
	c.JSON(200, obj)
}

func router(router *gin.Engine) {
	router.GET("/", myMiddleware1(), DefaultGet)
}

//中间件
func myMiddleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("I am the middleware")
		c.Next()
		fmt.Println("I am the after middleware")
	}
}

type Item struct {
	Id     int64
	ItemNo string `xorm:"item_no varchar(255) notnull"`
}

func (i Item) TableName() string {
	return "items"
}
