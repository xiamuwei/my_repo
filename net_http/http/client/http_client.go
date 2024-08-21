package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	// 
	resp, err := http.Get("http://localhost:3001")
	if err != nil {
		fmt.Println("client Get error = ", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        fmt.Println("Error reading response body:", err.Error())
        return
    }

    fmt.Println(string(body)) 
}