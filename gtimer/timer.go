package gtimer

import "time"

const (
	HourName = "HOUR"
	// HourInterval 小时间隔ms为精度
	HourInterval = 60 * 60 * 1e3
	// HourScales  12小时制
	HourScales = 12

	MinuteName = "MINUTE"
	// MinuteInterval 每分钟时间间隔
	MinuteInterval = 60 * 1e3
	// MinuteScales 60分钟
	MinuteScales = 60

	SecondName = "SECOND"
	// SecondInterval 秒的间隔
	SecondInterval = 1e3
	// SecondScales  60秒
	SecondScales = 60

	// TimersMaxCap //每个时间轮刻度挂载定时器的最大个数
	TimersMaxCap = 2048
)

type Timer struct {
	delayFuncation *DelayFunc
	// 调用时间 ms
	calltime int64
}

func UnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

func NewTimer(df *DelayFunc, nano int64) *Timer {
	return &Timer{
		delayFuncation: df,
		calltime:       nano / 1e6,
	}
}

func NewTimerAfter(df *DelayFunc, duation time.Duration) *Timer {
	return NewTimer(df, time.Now().UnixNano()+int64(duation))
}

func (t *Timer) Run() {
	go func() {
		now := UnixMilli()
		if t.calltime > now {
			time.Sleep(time.Duration(t.calltime-now) * time.Millisecond)
		}
		t.delayFuncation.Call()
	}()
}
