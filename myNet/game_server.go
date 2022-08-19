package myNet

import (
	"fmt"
	"leexsh/TCPGame/TCPGameServer/iface"
	"leexsh/TCPGame/TCPGameServer/utils"
	"net"
	"strconv"
)

/*
	game server模块 实现iserver借款==接口
*/

type GameServer struct {
	Name      string        // 服务名称
	IPVersion string        // IP版本
	IP        string        // IP地址
	Port      string        // 端口
	Router    iface.IRouter // 路由
}

func (g *GameServer) AddRouter(router iface.IRouter) {
	g.Router = router
	fmt.Println("[server]add router success")
}

func (g *GameServer) Start() {
	fmt.Printf("[Server]server name:%s is running, IP: %s, port:%s", g.Name, g.IP, g.Port)
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
			dealConn := NewConnection(conn, cid, g.Router)
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
	utils.LoadConfig()
	s := &GameServer{
		Name:      utils.YmlConfig.GlobalConfig.Name,
		IPVersion: "tcp4",
		IP:        utils.YmlConfig.GlobalConfig.IP,
		Port:      strconv.Itoa(utils.YmlConfig.GlobalConfig.TcpPort),
		Router:    nil,
	}
	return s
}
