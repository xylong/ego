package iface

// IRouter 路由接口
type IRouter interface {
	// 处理业务前钩子方法
	Before(IRequest)
	// 处理业务中钩子方法
	Handle(IRequest)
	// 处理业务后钩子方法
	After(IRequest)
}
