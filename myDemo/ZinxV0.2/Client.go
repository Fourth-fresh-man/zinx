package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)

	// 1、获取连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start error...Exit!")
		return
	}
	// 2、连接调用write写数据
	for {
		_, err := conn.Write([]byte("Hello Zinx v0.2"))
		if err != nil {
			fmt.Println("Write conn err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err", err)
			return
		}

		fmt.Printf("Server callback content=%s cnt=%d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}