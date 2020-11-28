package abstract

// Server 服务
type Server interface {
	// Start 启动服务
	Start()
	// Stop 停止服务
	Stop()
	// Run 运行服务
	Run()
}
