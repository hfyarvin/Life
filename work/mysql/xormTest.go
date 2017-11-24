package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:12345678@tcp(127.0.0.1:3306)/maxiiot_test?charset=utf8")
	checkErr(err)
	pintErr := engine.Ping()
	checkErr(pintErr)
	// engine.Close()
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	r := gin.Default()
	r.GET("/item", func(c *gin.Context) {
		//执行数据库语句
		// rs, err := engine.Query("select * from offices")

		// 获取单条
		item := new(User)
		_, err := engine.Where("id > ?", 1).Get(item)
		// var id int64
		// _, err := engine.Table("items").Where("item_no = ?", "A-1").Get(&id)
		//获取多条-------------------------------------------------
		// var List []Item
		// err := engine.Where("id > ?", 1).Find(&List)
		// var idList []int64
		// err := engine.Table("items").Where("id = ?", 1).Cols("id").Find(&idList)
		if err != nil {
			fmt.Println("find err :", err)
			c.JSON(200, "hello world")
		} else {
			c.JSON(200, item)
		}
	})
	r.Run(":8888")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Item struct {
	Id     int64
	ItemNo string `xorm:"item_no varchar(255) notnull"`
}

func (i Item) TableName() string {
	return "items"
}

type User struct {
	Id   int64
	Name string `xorm:"name varchar(255) notnull"`
	Sex  int64  `xorm:"sex int(11)"`
}

func (table User) TableName() string {
	return "maxiiot_users"
}
