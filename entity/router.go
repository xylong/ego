package entity

import "ego/iface"

// BaseRouter 基础路由
type BaseRouter struct{}

func (r *BaseRouter) Before(request iface.IRequest) {}
func (r *BaseRouter) Handle(request iface.IRequest) {}
func (r *BaseRouter) After(request iface.IRequest)  {}
