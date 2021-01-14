package entity

import (
	"ego/iface"
	"fmt"
	"log"
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
	handleAPI iface.HandleFunc
	// 告知当前连接已经退出/停止的channel
	ExitChan chan bool
}

// NewConnection 创建连接
func NewConnection(conn *net.TCPConn, connID uint32, handleAPI iface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: handleAPI,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
}

// Start 启动连接
func (c *Connection) Start() {
	go c.StartReader()
}

// StartReader 启动读数据业务
func (c *Connection) StartReader() {
	fmt.Println("reader is running")
	defer fmt.Printf("connID=%d reader exit,remoteAddr:%s", c.ConnID, c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端数据到buf中，最大512字节
		buf := make([]byte, 512)
		length, err := c.Conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			continue
		}

		// 调用绑定的handleAPI
		if err := c.handleAPI(c.Conn, buf, length); err != nil {
			fmt.Printf("connID=%d has error:%s", c.ConnID, err.Error())
			break
		}
	}
}

// Stop 停止连接
func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	c.isClosed = true
	if err := c.Conn.Close(); err != nil {
		log.Fatal(err)
	}
	close(c.ExitChan)
}

// GetTCPConnection 获取当前连接
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// GetRemoteAddr 获取远程地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送
func (c *Connection) Send(bytes []byte) error {
	return nil
}
