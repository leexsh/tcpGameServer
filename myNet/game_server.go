package myNet

import (
	"errors"
	"fmt"
	"net"
)

type GameServer struct {
	Name      string // 服务名称
	IPVersion string // IP版本
	IP        string // IP地址
	Port      string // 端口
}

func CallBackToClient(conn *net.TCPConn, buf []byte, cnt int) error {
	_, err := conn.Write(buf[:cnt])
	if err != nil {
		return errors.New("call back fail")
	}
	return nil
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
		var cid uint32 = 0
		// 3.listen
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				continue
			}
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			go dealConn.Start()
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
