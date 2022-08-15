package myNet

import "leexsh/TCPGame/TCPGameServer/iface"

type Request struct {
	// 客户端的连接
	Conn iface.IConnection
	// 客户端请求的数据
	Data []byte
}

func (r *Request) GetConnection() iface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Data
}
