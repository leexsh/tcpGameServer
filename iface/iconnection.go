package iface

import "net"

/*
	connect连接的抽象
*/

type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取conn
	GetTCPConnnecetion() *net.TCPConn
	// 获取conn id
	GetConnId() uint32
	// 获取对端ip:端口
	RemoteAddr() net.Addr
	// 发送数据
	SendMsg(uint32, uint32, []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
