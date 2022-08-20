package myNet

import (
	"TCPGameServer/iface"
	"TCPGameServer/utils"
	"errors"
	"fmt"
	"io"
	"net"
)

/*
	connect连接
*/

type Connection struct {
	Conn *net.TCPConn

	// conn id
	ConnID uint32
	// bool is close
	IsClosed bool

	ExitChan   chan bool
	MsgHandler iface.IMessageHandler
	MsgChan    chan []byte // 读写协程之间的chan(无缓冲)
}

func NewConnection(conn *net.TCPConn, id uint32, handler iface.IMessageHandler) *Connection {
	utils.LoadConfig()
	c := &Connection{
		Conn:       conn,
		ConnID:     id,
		ExitChan:   make(chan bool, 1),
		IsClosed:   false,
		MsgHandler: handler,
		MsgChan:    make(chan []byte),
	}
	return c
}

// read Goroutine
func (c *Connection) StartRead() {
	fmt.Println("read Goroutine is running")
	defer fmt.Println("read Goroutine has existed")
	// when closed, send exit flag to write Goroutine
	defer c.Stop()
	for {
		// 1.read head 8 bytes
		headData := make([]byte, DataPackTool.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnnecetion(), headData); err != nil {
			fmt.Println("read error")
			break
		}
		// 2. unpack
		msg, err := DataPackTool.UnPack(headData)
		if err != nil {
			break
		}
		// 3. read data
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnnecetion(), data); err != nil {
				fmt.Println("[server]Read data error")
				break
			}
		}
		msg.SetData(data)
		req := &Request{
			conn: c,
			msg:  msg,
		}

		// 由对应的msgId对应的router进行处理
		go c.MsgHandler.DoMsgHandler(req)
	}
}

// write Goroutine
func (c *Connection) StartWrite() {
	fmt.Println("writer Goroutine is running")
	defer fmt.Println("writer Goroutine has existed")
	for {
		select {
		case data := <-c.MsgChan:
			_, err := c.Conn.Write(data)
			if err != nil {
				fmt.Println("send to client error")
				return
			}
		case <-c.ExitChan:
			// reader exit, so notify writer to exit
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("[server] conn is start, conn id is:", c.ConnID)
	go c.StartRead()
	go c.StartWrite()
}

func (c *Connection) Stop() {
	fmt.Println("[server]conn close, conn id is: ", c.ConnID)
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	defer c.Conn.Close()
	// exit
	c.ExitChan <- true
	close(c.ExitChan)
	close(c.MsgChan)
}

func (m *Message) SendMsg() {

}

func (c *Connection) GetTCPConnnecetion() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgId uint32, msgType uint32, data []byte) error {
	if c.IsClosed {
		return errors.New("[server]conn is closed")
	}

	// pack
	binaryData, err := DataPackTool.Pack(NewMsg(msgId, msgType, data))
	if err != nil {
		return err
	}

	// send to write Goroutine
	c.MsgChan <- binaryData
	return nil
}
