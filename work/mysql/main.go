package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// var Db *sql.DB

func main() {
	DbInit()
}

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = "12345678"       //密码
	dbname     = "yiibaidb"       //表名
)

func DbInit() {
	// var err error
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbusername, dbpassword, dbhostsip, dbname)
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS yiibaidb.hello(world varchar(50))")
	if err != nil {
		log.Fatalln(err)
	}
	rows, err := db.Query("SELECT customerName FROM yiibaidb.customers")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("found row containing %q", s)
	}
	rows.Close()

}

func addNewData(db *sql.DB) (sql.Result, error) {
	rs, err := db.Exec("INSERT INTO test.hello(world) VALUES ('hello world')")
	return rs, err
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
