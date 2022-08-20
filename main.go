package main

import (
	"fmt"
	"leexsh/TCPGame/TCPGameServer/iface"
	"leexsh/TCPGame/TCPGameServer/myNet"
)

type PingRouter struct {
	myNet.BaseRouter
}

func (p *PingRouter) Handle(req iface.IReqeust) {
	fmt.Println("[server]Call Router Handle")
	fmt.Println("[server]receive from client, msg id is:", req.GetDataId(), " data is: ", string(req.GetData()))
	err := req.GetConnection().SendMsg(1, []byte("ping, ping"))
	if err != nil {
		fmt.Println("[server]err in router")
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
