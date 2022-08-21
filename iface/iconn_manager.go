package iface

/*
	连接管理抽象层
*/

type IConnManager interface {
	Add(connection IConnection)
	Remove(connection IConnection)
	Get(uint32) IConnection
	Len() int
	ClearConn()
}
