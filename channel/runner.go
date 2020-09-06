package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// 定义一个定时器
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	task      []func(int)
}

var ErrTimeOunt = errors.New("timeout")
var ErrInterrupt = errors.New("interrupt")

func NewRunner(timeout time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(timeout), // time.After()执行完成后，返回一个<-chan Time，所以才会有倒计时的效果
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.task = append(r.task, tasks...)
}

func (r *Runner) getInterrupt() bool {
	// 经典select用法，当执行到select时，若存在default，则执行通过，不再阻塞等待
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) run() error {
	for id, t := range r.task {
		if r.getInterrupt() {
			return ErrInterrupt
		}

		t(id)
	}

	return nil
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	// select阻塞等待
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOunt
	}
}
