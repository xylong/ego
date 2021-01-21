package main

import (
	"ego/entity"
	"ego/iface"
	"fmt"
)

type PingRouter struct {
	entity.BaseRouter
}

func (r *PingRouter) Before(request iface.IRequest) {
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("before...\n")); err != nil {
		fmt.Println("callback before error")
	}
}

func (r *PingRouter) Handle(request iface.IRequest) {
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte(fmt.Sprintf("%s\n", string(request.GetData())))); err != nil {
		fmt.Println("callback handle error")
	}
}

func (r *PingRouter) After(request iface.IRequest) {
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("after...\n")); err != nil {
		fmt.Println("callback after error")
	}
}

func main() {
	server := entity.NewServer("test")
	server.AddRouter(&PingRouter{})
	server.Run()
}
