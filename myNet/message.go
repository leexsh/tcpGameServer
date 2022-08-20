package myNet

/*
	消息的封装
*/

type Message struct {
	ID      uint32 // 消息ID
	Type    uint32 // 消息类型
	DataLen uint32 // 消息长度
	Data    []byte // 消息内容
}

func NewMsg(msgId, msgType uint32, data []byte) *Message {
	return &Message{
		ID:      msgId,
		DataLen: uint32(len(data)),
		Type:    msgType,
		Data:    data,
	}
}

func (m *Message) SetMsgType(u uint32) {
	m.Type = u
}

func (m *Message) GetMsgType() uint32 {
	return m.Type
}

func (m *Message) SetMsgID(u uint32) {
	m.ID = u
}

func (m *Message) SetDataLen(u uint32) {
	m.DataLen = u
}

func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}

func (m *Message) GetMsgId() uint32 {
	return m.ID
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}
