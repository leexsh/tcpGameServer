package test

import (
	"TCPGameServer/gtimer"
	"fmt"
	"testing"
)

func mytest(msg ...interface{}) {
	fmt.Println("msg is : ", msg[0].(string))
}

func TestDelayFunc(t *testing.T) {
	df := gtimer.NewDelayFunc(mytest, []interface{}{"hello"})
	df.Call()
}
