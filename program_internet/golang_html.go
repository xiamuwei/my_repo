package main 

import (
	"net/http"
	"fmt"

	"golang_html/method"
)

func main() {
	// 注册handle
	http.HandleFunc("/", method.Begin_html)
	http.HandleFunc("/login", method.Login_html)
	// 监听并服务
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("login server error = ", err.Error())
		return
	}
}
