package iface

/*

	消息管理的抽象层
*/

type IMessageHandler interface {
	// 执行对应的router方法
	DoMsgHandler(request IReqeust) error

	// add router
	AddRouter(msgTpye uint32, router IRouter) error
}
