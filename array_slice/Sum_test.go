package array_slice

import (
	"reflect"
	"testing"
)

/**
测试覆盖率工具：
	通过运行go test -cover Sum_test.go Sum.go  可以进行测试覆盖率计算
*/
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d, but want %d, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumBySlice(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d, but want %d, %v", got, want, numbers)
		}
	})
}

//func TestSum2(t *testing.T) {
//	numbers := [4]int{1, 2, 3, 4}
//
//	// 报错 数组长度也是类型的一部分：cannot use numbers (type [4]int) as type [5]int in argument to Sum
//	got := Sum(numbers)
//	want := 10
//	if got != want {
//		t.Errorf("got %d, but want %d, %v", got, want, numbers)
//	}
//}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// golang不能对切片使用等号运算符，切片只能和nil做比较
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	// 由于checkSums函数限制了got、want的类型，所以传入别的类型时，会编译期报错
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
