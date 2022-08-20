package myNet

import (
	"TCPGameServer/iface"
	"TCPGameServer/utils"
	"fmt"
	"net"
	"strconv"
)

/*
	game server模块 实现iserver借款==接口
*/

type Server struct {
	Name      string                // 服务名称
	IPVersion string                // IP版本
	IP        string                // IP地址
	Port      string                // 端口
	MsgHander iface.IMessageHandler // 路由
}

func (g *Server) AddRouter(msgType uint32, router iface.IRouter) {
	g.MsgHander.AddRouter(msgType, router)
	fmt.Println("[server]add router success")
}

func (g *Server) Start() {
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
			dealConn := NewConnection(conn, cid, g.MsgHander)
			cid++
			go dealConn.Start()
		}
	}()

}

func (g *Server) Stop() {
	// todo: stop
}

func (g *Server) Serve() {
	g.Start()

	// do another things

	// block
	select {}
}

func NewServer(name string) *Server {
	utils.LoadConfig()
	return &Server{
		Name:      utils.YmlConfig.GlobalConfig.Name,
		IPVersion: "tcp4",
		IP:        utils.YmlConfig.GlobalConfig.IP,
		Port:      strconv.Itoa(utils.YmlConfig.GlobalConfig.TcpPort),
		MsgHander: NewMsgHandle(),
	}
}
