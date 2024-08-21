package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册回调函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "this is a test for golang http")
	})

	// 监听端口
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("http ListenAndServe error = ", err.Error())
	}

}