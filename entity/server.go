package entity

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
}

// Start 启动
func (s *Server) Start() {

}

// Stop 停止
func (s *Server) Stop() {

}

// Run 运行
func (s Server) Run() {

}