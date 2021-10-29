package znet

import (
	"zinx/ziface"
)
// BaseRouter 先实现所有方法，作为基类，将来可以有选择性去实现即可
type BaseRouter struct {}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {}

func (b *BaseRouter) Handle(request ziface.IRequest) {}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {}


