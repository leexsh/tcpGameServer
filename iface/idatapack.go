package iface

/*
	TCP数据的封包和拆包
*/

type IDataPack interface {
	// 获取包头长度
	GetHeadLen() uint32

	// 封包
	Pack(msg IMessage) ([]byte, error)

	// 拆包
	UnPack([]byte) (IMessage, error)
}
