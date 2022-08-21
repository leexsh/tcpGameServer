package iface

/*
	抽象的server接口
*/

type IServer interface {
	// 启动服务
	Start()
	// 停止服务
	Stop()
	// 运行
	Serve()

	AddRouter(uint32, IRouter)

	GetConnManager() IConnManager

	// register start hook method
	SetOnConnStart(func(conn IConnection))
	// register stop hook method
	SetOnConnStop(func(conn IConnection))

	// call start hook
	CallOnConnStart(connection IConnection)
	// call stop hook
	CallOnCOnnStop(connection IConnection)

	Packet() IDataPack
}
