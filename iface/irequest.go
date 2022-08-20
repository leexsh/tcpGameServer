package iface

/*
	抽象的请求
*/

type IReqeust interface {
	// 得到当前请求
	GetConnection() IConnection
	// 得到当前请求的数据
	GetData() []byte
	GetDataId() uint32
	GetMsgType() uint32
}
