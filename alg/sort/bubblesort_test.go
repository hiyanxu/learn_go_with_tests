package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort_Sort(t *testing.T) {
	a := []int{5, 9, 1, 4, 7}
	fmt.Println(new(bubbleSort).Sort(a))
}
