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

	AddRouter(router IRouter)
}
