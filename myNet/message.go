package myNet

/*
	消息的封装
*/

type Message struct {
	ID      uint32 // 消息ID
	DateLen uint32 // 消息长度
	Data    []byte // 消息内容
}

func (m *Message) SetMsgID(u uint32) {
	m.ID = u
}

func (m *Message) SetMsgLen(u uint32) {
	m.DateLen = u
}

func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}

func (m *Message) GetMsgId() uint32 {
	return m.ID
}

func (m *Message) GetMsgLen() uint32 {
	return m.DateLen
}

func (m *Message) GetData() []byte {
	return m.Data
}
