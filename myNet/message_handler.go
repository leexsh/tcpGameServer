package myNet

import (
	"TCPGameServer/iface"
	"TCPGameServer/utils"
	"errors"
	"fmt"
)

/*
	消息处理模块的实现
*/

type MessageHandle struct {
	// router 集合
	Apis map[uint32]iface.IRouter

	// worker的消息队列
	TaskQueue []chan iface.IReqeust
	// 工作池的worker的数量
	WorkPoolSize uint32
}

func NewMsgHandle() *MessageHandle {
	utils.LoadConfig()
	return &MessageHandle{
		Apis:         make(map[uint32]iface.IRouter, 32),
		TaskQueue:    make([]chan iface.IReqeust, utils.YmlConfig.GlobalConfig.WorkPoolSize),
		WorkPoolSize: utils.YmlConfig.GlobalConfig.WorkPoolSize, // 全局配置获取
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

func (m *MessageHandle) StartWorkPool() {
	// 开启worker pool
	for i := 0; i < int(m.WorkPoolSize); i++ {
		m.TaskQueue[i] = make(chan iface.IReqeust, utils.YmlConfig.GlobalConfig.MaxWorkPoolSize)
		go m.startOneWokrer(i, m.TaskQueue[i])
	}
}

func (m *MessageHandle) startOneWokrer(workerID int, taskQueue chan iface.IReqeust) {
	fmt.Println("start worker id: ", workerID)
	// ，每个worker都在堵塞等待队列中来的req
	for true {
		select {
		case req := <-taskQueue:
			m.DoMsgHandler(req)
		}
	}
}

func (m *MessageHandle) DispatchMsg(req iface.IReqeust) {
	// 轮训 将消息发送给worker
	workerId := req.GetDataId() % m.WorkPoolSize
	fmt.Println("worker pool id is: ", workerId)
	m.TaskQueue[workerId] <- req
}
