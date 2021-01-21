package entity

import (
	"ego/iface"
	"fmt"
	"log"
	"net"
)

// Connection 连接
type Connection struct {
	// 当前连接的socket
	Conn *net.TCPConn
	// 连接id
	ConnID uint32
	// 当前连接状态
	isClosed bool
	// 告知当前连接已经退出/停止的channel
	ExitChan chan bool
	// 该链接处理的方法路由
	Router iface.IRouter
}

// NewConnection 创建连接
func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
}

// Start 启动连接
func (c *Connection) Start() {
	fmt.Printf("conn:%d start", c.ConnID)
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
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 执行注册的路由方法
		go func(request iface.IRequest) {
			c.Router.Before(request)
			c.Router.Handle(request)
			c.Router.After(request)
		}(&Request{
			conn: c,
			data: buf,
		})
	}
}

// Stop 停止连接
func (c *Connection) Stop() {
	fmt.Printf("conn:%d stop", c.ConnID)

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
