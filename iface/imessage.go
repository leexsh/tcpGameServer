package iface

/*
	抽象的消息结构体封装
*/
type IMessage interface {
	SetMsgID(uint32)
	SetDataLen(uint32)
	SetData([]byte)

	GetMsgId() uint32
	GetDataLen() uint32
	GetData() []byte
}
