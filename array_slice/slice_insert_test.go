package array_slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppendNormal(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3, 4}
	s1 = AppendNormal(s1, 4)
	if reflect.DeepEqual(s1, s2) {
		fmt.Println(s1)
	} else {
		t.Errorf("fail: %d", s2)
	}
}
