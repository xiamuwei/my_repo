package main

import (
	"net"
	"fmt"
)

func main() {
	addr := "localhost:8001"
	// 监听端口
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("error = ", err.Error())
	}
	defer ln.Close()

	// 循环接受请求
	for {
		conn ,err := ln.Accept()
		if err != nil {
			fmt.Println("request error = ", err.Error())
			continue
		}
		// 回应请求
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 128)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("handleConnection error = ", err.Error())
			return
		}
		fmt.Println(string(buf[:n]))
	}
}