package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

// 方便测试，注入一个io.Writer的接口，模拟向不同writer中写入数据的情况
type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func main() {
	// 通过在main函数中注入真正的sleeper执行time.Sleep()
	// 而在测试时，简单地通过一个计数器模仿实现sleep
	//spySleeper := &SpySleeper{Calls: 0}
	o := ConfigurableSleeper{duration: 1}
	Countdown(os.Stdout, &o)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	// fmt.Fprintf函数用于将对应的字符串写入到前面的writer中，至于writer是标准输出，还是buffer，都可以
	//fmt.Fprintf(writer, "3")
	for i := countdownStart; i > 0; i-- {
		//time.Sleep(1 * time.Second)
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}
