package myNet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"leexsh/TCPGame/TCPGameServer/iface"
	"leexsh/TCPGame/TCPGameServer/utils"
)

/*
	TCP数据流的封包和拆包
*/

type DataPack struct {
}

var DataPackTool = &DataPack{}

func (d *DataPack) GetHeadLen() uint32 {
	// uint32 4 byte so head is 8 byte
	return 8
}

func (d *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	// 1. create buffer
	buf := bytes.NewBuffer([]byte{})
	// 2. write data len
	err := binary.Write(buf, binary.LittleEndian, msg.GetDataLen())
	if err != nil {
		return nil, err
	}
	// 3. write data id
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 4. write data
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *DataPack) UnPack(bytesData []byte) (iface.IMessage, error) {
	dataBuf := bytes.NewReader(bytesData)
	msg := &Message{}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	if utils.YmlConfig.GlobalConfig.MaxPackageSize > 0 &&
		msg.GetDataLen() > utils.YmlConfig.GlobalConfig.MaxPackageSize {
		return nil, errors.New("package too large")
	}

	return msg, nil
}
