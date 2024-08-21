package main

import (
	"net"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	// 拨号
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		fmt.Println("client error = ", err.Error())
	}
	defer conn.Close()

	// 发送请求
	// str := "golang can build tcp connection"
	// _, err = conn.Write([]byte(str))
	// if err != nil {
	// 	fmt.Println("client send message error = ", err.Error())
	// }
	dir, _ := os.Getwd()
	fmt.Println("dir = ", dir)
	byteBuf, err1 := ioutil.ReadFile("../../read.txt")
	if err1 != nil {
		fmt.Println("client byteBuf error = ", err.Error())
		return
	} 
	_, err1 = conn.Write(byteBuf)
	if err1 != nil {
		fmt.Println("client write error = ", err.Error())
		return
	} 
}