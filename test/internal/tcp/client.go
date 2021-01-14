package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 1.连接服务
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatal(err)
	}

	// 2. 发消息
	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		buf := make([]byte, 512)
		length, err := conn.Read(buf)
		fmt.Println(string(buf[:length]))

		time.Sleep(time.Second * 2)
	}
}
