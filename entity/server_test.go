package entity

import (
	"net"
	"testing"
	"time"
)

// TestServer_Run 测试服务运行
func TestServer_Run(t *testing.T) {
	// 启动服务
	go func() {
		server := NewServer("test")
		server.Run()
	}()
	// 启动客户端
	<-send(t)
}

func send(t *testing.T) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer func() {
			done <- struct{}{}
		}()
		time.Sleep(time.Second * 1) // 等服务器启动

		index := 0
		conn, err := net.Dial("tcp", "127.0.0.1:10000")
		if err != nil {
			t.Error(err)
			return
		}

		for {
			_, err := conn.Write([]byte("hello"))
			if err != nil {
				t.Error(err)
				return
			}

			buf := make([]byte, 512)
			length, err := conn.Read(buf)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(string(buf[:length]))

			time.Sleep(time.Second * 2)
			index++
			if index > 2 {
				break
			}
		}
	}()

	return done
}
