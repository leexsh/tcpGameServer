package iface

type IServer interface {
	// 启动服务
	Start()
	// 停止服务
	Stop()
	// 运行
	Serve()
}
