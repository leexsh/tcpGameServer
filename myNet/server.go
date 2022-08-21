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
	Name        string                // 服务名称
	IPVersion   string                // IP版本
	IP          string                // IP地址
	Port        string                // 端口
	MsgHandler  iface.IMessageHandler // 路由
	ConnManager iface.IConnManager    // conn manager
}

func (g *Server) AddRouter(msgType uint32, router iface.IRouter) {
	g.MsgHandler.AddRouter(msgType, router)
}

func (g *Server) Start() {
	fmt.Printf("[Server]server name:%s is running, IP: %s, port:%s\n", g.Name, g.IP, g.Port)
	go func() {
		// 0.start work pool
		g.MsgHandler.StartWorkPool()

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
			// 最大连接个数判断
			if g.ConnManager.Len() > utils.YmlConfig.GlobalConfig.MaxConn {
				conn.Write([]byte("connection limit"))
				conn.Close()
				continue
			}

			dealConn := NewConnection(g, conn, cid, g.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()

}

func (g *Server) Stop() {
	fmt.Println("server stop ")
	g.ConnManager.ClearConn()
	g.MsgHandler.StopWorkPool()
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
		Name:        utils.YmlConfig.GlobalConfig.Name,
		IPVersion:   "tcp4",
		IP:          utils.YmlConfig.GlobalConfig.IP,
		Port:        strconv.Itoa(utils.YmlConfig.GlobalConfig.TcpPort),
		MsgHandler:  NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
}

func (g *Server) GetConnManager() iface.IConnManager {
	return g.ConnManager
}
