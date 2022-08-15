package main

import (
	"fmt"
	"leexsh/TCPGame/TCPGameServer/iface"
	"leexsh/TCPGame/TCPGameServer/myNet"
)

type PingRouter struct {
	myNet.BaseRouter
}

func (p *PingRouter) PreHandle(req iface.IReqeust) {
	fmt.Println("[server]Call Router Prehandle")
	_, err := req.GetConnection().GetTCPConnnecetion().Write([]byte("before ping\n"))
	if err != nil {
		fmt.Println("err before router")
		return
	}
}

func (p *PingRouter) Handle(req iface.IReqeust) {
	fmt.Println("[server]Call Router Handle")
	_, err := req.GetConnection().GetTCPConnnecetion().Write([]byte("ping\n"))
	if err != nil {
		fmt.Println("err in router")
		return
	}
}

func (p *PingRouter) AfterHandle(req iface.IReqeust) {
	fmt.Println("[server]Call Router Afterhandle")
	_, err := req.GetConnection().GetTCPConnnecetion().Write([]byte("after\n"))
	if err != nil {
		fmt.Println("err after  router")
		return
	}
}

func main() {
	s := myNet.NewServer("server")
	// add router
	s.AddRouter(&PingRouter{})
	s.Start()
	select {}

}
