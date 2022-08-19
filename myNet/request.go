package myNet

import "leexsh/TCPGame/TCPGameServer/iface"

/*
	客户端请求的封装
*/

type Request struct {
	// 客户端的连接
	conn iface.IConnection
	// 客户端请求的数据
	msg iface.IMessage
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetDataId() uint32 {
	return r.msg.GetMsgId()
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}
