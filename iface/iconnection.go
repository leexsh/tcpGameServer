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
	// 设置连接属性
	SetProperty(key string, value interface{})
	// 获取链接属性
	GetProperty(key string) (interface{}, error)
	// 移除连接属性
	RemoveProperty(key string)
}

type HandleFunc func(*net.TCPConn, []byte, int) error
