package test

import (
	"TCPGameServer/myNet"
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataPack(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		t.Fatal(err)
	}

	// server
	go func() {
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal(err)
		}

		go func(conn net.Conn) {

			for {
				// 1. read head
				headData := make([]byte, myNet.DataPackTool.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					t.Fatal(err)
				}
				msgHead, err := myNet.DataPackTool.UnPack(headData)
				if err != nil {
					t.Fatal(err)
				}
				if msgHead.GetDataLen() > 0 {
					// 2. read data
					msg := msgHead.(*myNet.Message)
					msg.Data = make([]byte, msg.GetDataLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						t.Fatal(err)
					}
					fmt.Println("id:", msg.ID, ", len: ", msg.DataLen, " type is:", msg.Type, " data: ", string(msg.Data))
				}

			}
		}(conn)
	}()

	// client
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		t.Fatal(err)
	}
	msg1 := &myNet.Message{
		ID:      1,
		DataLen: 4,
		Type:    1,
		Data:    []byte{'1', '2', '3', '4'},
	}
	data1, err := myNet.DataPackTool.Pack(msg1)
	if err != nil {
		t.Fatal(err)
	}
	msg2 := &myNet.Message{
		ID:      2,
		DataLen: 5,
		Type:    1,
		Data:    []byte{'h', 'e', 'l', 'l', '0'},
	}
	data2, err := myNet.DataPackTool.Pack(msg2)

	// data1 data2 放在一起
	data1 = append(data1, data2...)
	// send
	conn.Write(data1)

	time.Sleep(time.Second * 2)
}
