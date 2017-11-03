package main

import (
	"fmt"
	"io/ioutil"
	// "github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	resp, _ := http.Get("http://0.0.0.0:8888/test1")
	helpRead(resp)

	//post多传两个参数
	resp, _ = http.Post("http://0.0.0.0:8888/test2", "", strings.NewReader(""))
	helpRead(resp)

	//PostForm
	resp, _ = http.Post("http://0.0.0.0:8888/postform", "application/x-www-form-urlencoded", strings.NewReader("message=8888888&extra=999999"))
	helpRead(resp)
}

func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}

	fmt.Println(string(body))
}
