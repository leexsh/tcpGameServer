package test

import (
	"fmt"
	"io"
	"leexsh/TCPGame/TCPGameServer/myNet"
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
		dp := myNet.NewDataPack()
		go func(conn net.Conn, pack *myNet.DataPack) {

			for {
				// 1. read head
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					t.Fatal(err)
				}
				msgHead, err := dp.UnPack(headData)
				if err != nil {
					t.Fatal(err)
				}
				if msgHead.GetMsgLen() > 0 {
					// 2. read data
					msg := msgHead.(*myNet.Message)
					msg.Data = make([]byte, msg.GetMsgLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						t.Fatal(err)
					}
					fmt.Println("id:", msg.ID, ", len: ", msg.DateLen, " data: ", string(msg.Data))
				}

			}
		}(conn, dp)
	}()

	// client
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		t.Fatal(err)
	}
	dp := myNet.NewDataPack()
	msg1 := &myNet.Message{
		ID:      1,
		DateLen: 4,
		Data:    []byte{'1', '2', '3', '4'},
	}
	data1, err := dp.Pack(msg1)
	if err != nil {
		t.Fatal(err)
	}
	msg2 := &myNet.Message{
		ID:      2,
		DateLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', '0'},
	}
	data2, err := dp.Pack(msg2)

	// data1 data2 放在一起
	data1 = append(data1, data2...)
	// send
	conn.Write(data1)

	time.Sleep(time.Second * 2)
}
