package net

import (
	"fmt"
	"net"
)

type GameServer struct {
	Name      string // 服务名称
	IPVersion string // IP版本
	IP        string // IP地址
	Port      string // 端口
}

func (g *GameServer) Start() {
	fmt.Printf("[Server]server is running, IP: %s, port:%s", g.IP, g.Port)
	go func() {
		// 1.get tcp addr
		addr, err := net.ResolveTCPAddr(g.IPVersion, fmt.Sprintf("%s:%s", g.IP, g.Port))
		if err != nil {
			return
		}
		// 2.create listen addr
		listener, err := net.ListenTCP(g.IPVersion, addr)
		if err != nil {
			return
		}
		fmt.Println("[Server]start server success")

		// 3.listen
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				continue
			}
			// go
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						continue
					}
					if _, err = conn.Write(buf[:cnt]); err != nil {
						fmt.Printf("write err")
						break
					}
				}
			}()
		}
	}()

}

func (g *GameServer) Stop() {
	// todo: stop
}

func (g *GameServer) Serve() {
	g.Start()

	// do another things

	// block
	select {}
}

func NewServer(name string) *GameServer {
	s := &GameServer{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      "8888",
	}
	return s
}
