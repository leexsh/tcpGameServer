package iface

type IRouter interface {
	// 处理conn前
	PreHandle(req IReqeust)
	// 处理conn
	Handle(req IReqeust)
	// 处理conn后
	AfterHandle(req IReqeust)
}
