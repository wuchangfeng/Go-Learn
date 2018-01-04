package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the sever...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // 程序终止
	}

	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

// 开启一个协程去处理请求
func doServerStuff(conn net.Conn) {
	for {
		// 使用一个512字节的缓冲data来读取客户端发送的内容
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
