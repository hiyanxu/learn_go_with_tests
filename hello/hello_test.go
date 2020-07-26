package main

import (
	"testing"
)

// 通过go test -test.run TestXXX  指定测试的函数

func TestHello(t *testing.T) {
	got := Hello3("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%q'  want '%q'", got, want)
	}
}

// t.Run()可以对同一个方法进行分组测试
func TestHello3(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello3("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("say hello when an empty string is supplied", func(t *testing.T) {
		got := Hello3("")
		want := "Hello, World"

		if got != want {
			t.Errorf("get %q, want %q", got, want)
		}
	})
}

func TestHello4(t *testing.T) {
	// go1.9新特性：t.Helper()需要告诉测试套件这个方法是辅助函数，当测试失败时报告的行数是函数调用的地方
	// 当t.Helper()注释时，报错： hello_test.go:42: get "Hello, World", want "Hello, World1"（非函数调用方）
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("get %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello4("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello4("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello4("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello4("hehe", "French")
		want := "Bonjour, hehe"
		assertCorrectMessage(t, got, want)
	})

}
