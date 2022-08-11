package main

import (
	"leexsh/tcpGameServer/net"
)

func main() {
	s := net.NewServer("myserver")
	s.Start()
	select {}

}
