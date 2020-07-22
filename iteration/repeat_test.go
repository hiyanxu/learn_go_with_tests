package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected: %q but got %q", expected, repeated)
	}
}

func TestRepeat1(t *testing.T) {
	repeated := Repeat1("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected: %q but got %q", expected, repeated)
	}
}

// 基准测试：通过测试CPU和内存的效率问题，来评估被测试代码的性能
/**
基准测试规则：
1、基准测试函数必须以Benchmark开头，必须是可导出的
2、参数必须是指向Benchmark类型的指针作为唯一参数
3、不能有返回值
4、被测试代码要放到for循环中

运行命令：
go test -bench=.  .表示运行所有的基准测试
*/
func BenchmarkRepeat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat1("a")
	}
}

func TestRepeat2(t *testing.T) {
	got := Repeat2("a", 4)
	want := "aaaa"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ExampleRepeat2() {
	got := Repeat2("b", 3)
	fmt.Println(got)
	// Output: bbb
}
