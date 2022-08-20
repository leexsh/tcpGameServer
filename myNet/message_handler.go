package myNet

import (
	"TCPGameServer/iface"
	"errors"
)

/*
	消息处理模块的实现
*/

type MessageHandle struct {
	Apis map[uint32]iface.IRouter
}

func NewMsgHandle() *MessageHandle {
	return &MessageHandle{
		make(map[uint32]iface.IRouter, 32),
	}
}

func (m *MessageHandle) DoMsgHandler(request iface.IReqeust) error {
	handler, ok := m.Apis[request.GetMsgType()]
	if !ok {
		return errors.New("not found this router")
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.AfterHandle(request)
	return nil
}

func (m *MessageHandle) AddRouter(msgType uint32, router iface.IRouter) error {
	if _, ok := m.Apis[msgType]; ok {
		return errors.New("[server]already exists this msg id")
	}
	m.Apis[msgType] = router
	return nil
}
