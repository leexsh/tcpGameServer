package myNet

import (
	"fmt"
	"leexsh/TCPGame/TCPGameServer/iface"
	"net"
)

type Connection struct {
	Conn *net.TCPConn

	// conn id
	ConnID uint32
	// bool is close
	IsClosed bool

	ExitChan chan bool
	Router   iface.IRouter
}

func NewConnection(conn *net.TCPConn, id uint32, router iface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   id,
		ExitChan: make(chan bool, 1),
		Router:   router,
		IsClosed: false,
	}
	return c
}

func (c *Connection) StartRead() {
	defer c.Conn.Close()
	for {
		// read form client
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("[server] read err, cnt is: ", cnt)
			break
		}
		// if err := c.HandleMethod(c.Conn, buf, cnt); err != nil {
		// 	break
		// }
		req := &Request{
			Conn: c,
			Data: buf,
		}
		// 从路由中找到对应的router
		go func(req iface.IReqeust) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.AfterHandle(req)
		}(req)
	}
}

func (c *Connection) Start() {
	fmt.Println("[server] conn is start, conn id is:", c.ConnID)
	c.StartRead()
}

func (c *Connection) Stop() {
	fmt.Println("[server]conn close, conn id is: ", c.ConnID)
	if c.IsClosed {
		return
	}
	c.IsClosed = true

	defer c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnnecetion() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
