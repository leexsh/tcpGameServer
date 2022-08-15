package myNet

import "leexsh/TCPGame/TCPGameServer/iface"

// BaseRouter 继承于IRouter  后续具体Router只需要适配BaseRouter即可
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(req iface.IReqeust) {

}

func (br *BaseRouter) Handle(req iface.IReqeust) {
}

func (br *BaseRouter) AfterHandle(req iface.IReqeust) {
}
