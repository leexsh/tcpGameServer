package test

import (
	"fmt"
	"net"
	"testing"
)

func TestConn(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 512)
	cnt, err := conn.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
	fmt.Println(cnt)
}
