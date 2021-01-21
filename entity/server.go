package entity

import (
	"ego/iface"
	"ego/util"
	"fmt"
	"net"
)

// Server 服务
type Server struct {
	// Name 服务器名称
	Name string
	// IPVersion IP版本
	IPVersion string
	// 监听IP
	IP string
	// Port 监听端口
	Port int
	// 路由
	Router iface.IRouter
}

// NewServer 创建server
func NewServer(name string) iface.IServer {
	return &Server{
		Name:      util.Configure.Name,
		IPVersion: "tcp4",
		IP:        util.Configure.Host,
		Port:      util.Configure.Port,
		Router:    nil,
	}
}

// Start 启动
func (s *Server) Start() {
	fmt.Printf("server listening at %s:%d\n", s.IP, s.Port)

	go func() {
		// 1.获取tcp 地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}
		// 2.监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("listen %s error:%s", s.IPVersion, err.Error())
			return
		}
		fmt.Println("start server ego successful...")
		// 3.阻塞等待客户端连接，处理业务
		var connID uint32
		connID = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept error:%s\n", err.Error())
				continue
			}

			// 调用连接处理逻辑
			connection := NewConnection(conn, connID, s.Router)
			go connection.Start()
			connID++
		}
	}()
}

// Stop 停止
func (s *Server) Stop() {

}

// Run 运行
func (s *Server) Run() {
	s.Start()

	// 阻塞
	select {}
}

// AddRouter 添加路由
func (s *Server) AddRouter(router iface.IRouter) {
	s.Router = router
}

// Answer 回复
func Answer(conn *net.TCPConn, bytes []byte, length int) error {
	_, err := conn.Write(bytes[:length])
	return err
}
