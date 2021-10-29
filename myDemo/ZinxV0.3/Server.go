package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// MyRouter 自定义的路由
type MyRouter struct {
	znet.BaseRouter
}

func (this *MyRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

func (this *MyRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

func (this *MyRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}


func main() {
	s := znet.NewServer("[Zinx v0.3]")
	s.AddRouter(&MyRouter{})
	s.Serve()
}


