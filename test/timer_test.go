package test

import (
	"TCPGameServer/gtimer"
	"fmt"
	"testing"
	"time"
)

func myFunc(v ...interface{}) {
	fmt.Println("hello")
}

func TestTimer(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func(i int) {
			gtimer.NewTimerAfter(gtimer.NewDelayFunc(myFunc, nil), time.Duration(2*i)*time.Second).Run()
		}(i)
	}
}
