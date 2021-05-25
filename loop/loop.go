package loop

import "time"

type Loop struct {
	isRunning bool

	interval time.Duration
	callback func()
}

func New(interval time.Duration, callback func()) *Loop {
	return &Loop{
		interval: interval,
		callback: callback,
	}
}

func (loop *Loop) run() {
	timeToCall := time.Now().Add(loop.interval)

	for loop.isRunning {
		timer := time.NewTimer(time.Until(timeToCall))
		<-timer.C
		loop.callback()
		timeToCall = timeToCall.Add(loop.interval)
	}
}

func (loop *Loop) Start() {
	loop.isRunning = true
	loop.run()
}

func (loop *Loop) Stop() {
	loop.isRunning = false
}
