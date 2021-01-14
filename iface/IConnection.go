package iface

import "net"

// IConnection 连接接口
type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接模块的连接id
	GetConnID() uint32
	// 获取远程客户端的tcp状态
	GetRemoteAddr() net.Addr
	// 发送数据
	Send(data []byte) error
}

// HandleFunc 处理连接业务函数
// *net.TCPConn 对端连接
// []byte 数据
// int	数据长度
type HandleFunc func(*net.TCPConn, []byte, int) error
