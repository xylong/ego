package entity

import (
	"ego/abstract"
	"net"
)

type Connection struct {
	// 当前连接的socket
	Conn *net.TCPConn
	// 连接id
	ConnID uint32
	// 当前连接状态
	isClosed bool
	// 当前连接绑定的处理业务方法API
	handleAPI abstract.HandleFunc
	// 告知当前连接已经退出/停止的channel
	ExitChan chan bool
}

// NewConnection 创建连接
func NewConnection(conn *net.TCPConn, connID uint32, handleAPI abstract.HandleFunc) *Connection {
	return &Connection{
		Conn: conn,
		ConnID: connID,
		handleAPI: handleAPI,
		isClosed: false,
		ExitChan: make(chan bool,1),
	}
}
