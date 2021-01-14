package iface

// IRequest 请求接口
type IRequest interface {
	// 获取连接
	GetConnection() IConnection
	// 获取数据
	GetData() []byte
}
