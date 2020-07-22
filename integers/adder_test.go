package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("expected %d, but got %d", expected, sum)
	}
}

// 示例 当增加Output，会在测试时自动执行该示例函数；若删除该注释，则函数会被编译，但不会执行
// 该示例代码会出现在doc文档中
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
