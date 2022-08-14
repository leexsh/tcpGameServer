package main

import (
	"leexsh/TCPGameServer/myNet"
)

func main() {
	s := myNet.NewServer("server")
	s.Start()
	select {}

}
