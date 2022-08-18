package iface

/*
	抽象的消息结构体封装
*/
type IMessage interface {
	SetMsgID(uint32)
	SetMsgLen(uint32)
	SetData([]byte)

	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte
}
