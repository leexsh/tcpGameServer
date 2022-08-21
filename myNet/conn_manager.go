package myNet

import (
	"TCPGameServer/iface"
	"fmt"
	"sync"
)

type ConnManager struct {
	// 连接集合
	conns map[uint32]iface.IConnection
	Lock  sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		conns: make(map[uint32]iface.IConnection),
	}
}

func (c *ConnManager) Add(connection iface.IConnection) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.conns[connection.GetConnId()] = connection
	fmt.Println("connect add success, conn id: ", connection.GetConnId())
}

func (c *ConnManager) Remove(connection iface.IConnection) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	delete(c.conns, connection.GetConnId())
}

func (c *ConnManager) Get(u uint32) iface.IConnection {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	if conn, ok := c.conns[u]; ok {
		return conn
	}
	return nil
}

func (c *ConnManager) Len() int {
	return len(c.conns)
}

func (c *ConnManager) ClearConn() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	for id, conn := range c.conns {
		conn.Stop()
		delete(c.conns, id)
	}
}

func (c *ConnManager) ClearOneConn(cid uint32) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	if conn, ok := c.conns[cid]; ok {
		conn.Stop()
		delete(c.conns, cid)
	}
}
