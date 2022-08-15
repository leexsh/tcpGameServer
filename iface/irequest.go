package iface

type IReqeust interface {
	// 得到当前请求
	GetConnection() IConnection
	// 得到当前请求的数据
	GetData() []byte
}
