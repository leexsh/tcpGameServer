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
	Name              string                // 服务名称
	IPVersion         string                // IP版本
	IP                string                // IP地址
	Port              string                // 端口
	MsgHandler        iface.IMessageHandler // 路由
	ConnManager       iface.IConnManager    // conn manager
	OnConnectionStart func(connection iface.IConnection)
	OnConnectionStop  func(connection iface.IConnection)
	packet            iface.IDataPack
	exitChan          chan struct{}
}

func (g *Server) AddRouter(msgType uint32, router iface.IRouter) {
	g.MsgHandler.AddRouter(msgType, router)
}

func (g *Server) Start() {
	fmt.Printf("[Server]server name:%s is running, IP: %s, port:%s\n", g.Name, g.IP, g.Port)
	g.exitChan = make(chan struct{})
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
		go func() {
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
		select {
		case <-g.exitChan:
			err := listener.Close()
			if err != nil {
				fmt.Println("listern err")
			}
		}

	}()
}

func (g *Server) Stop() {
	fmt.Println("server stop ")
	g.ConnManager.ClearConn()
	g.MsgHandler.StopWorkPool()
	g.exitChan <- struct{}{}
	close(g.exitChan)
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
		exitChan:    nil,
		packet:      DataPackTool,
	}
}

func (g *Server) GetConnManager() iface.IConnManager {
	return g.ConnManager
}

// hook methods
func (g *Server) SetOnConnStart(hookFunc func(conn iface.IConnection)) {
	g.OnConnectionStart = hookFunc
}

func (g *Server) SetOnConnStop(hookFunc func(conn iface.IConnection)) {
	g.OnConnectionStop = hookFunc
}

func (g *Server) CallOnConnStart(connection iface.IConnection) {
	if g.OnConnectionStart != nil {
		g.OnConnectionStart(connection)
	}
}

func (g *Server) CallOnCOnnStop(connection iface.IConnection) {
	if g.OnConnectionStop != nil {
		g.OnConnectionStop(connection)
	}
}

func (g *Server) Packet() iface.IDataPack {
	return g.packet
}

func init() {

}
