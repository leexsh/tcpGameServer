package gtimer

import "fmt"

/*
	延迟调用函数的设计
	定时器到了后，触发之前注册的函数
*/

type DelayFunc struct {
	delayFuncation func(...interface{}) // 延时函数
	args           []interface{}        // 延迟函数形参
}

func NewDelayFunc(f func(v ...interface{}), args []interface{}) *DelayFunc {
	return &DelayFunc{
		delayFuncation: f,
		args:           args,
	}
}

// call delay func
func (d *DelayFunc) Call() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error")
		}
	}()
	d.delayFuncation(d.args...)
}
