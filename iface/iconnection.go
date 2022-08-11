package iface

import "net"

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
	RemoteAddr() *net.Addr
	// 发送数据
	Send(data []byte) error
}
