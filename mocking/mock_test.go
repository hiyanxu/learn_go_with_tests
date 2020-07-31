package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{Calls: 0}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	//want := "3"
	// 反引号可以创建string，允许放置东西到新的一行
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}

	t.Run("sleep after every print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}
