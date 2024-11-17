package array_slice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Array() {
	s1 := []int{1, 2, 3}
	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Printf("s1, len: %d, cap: %d, slice header point: %p\n", len(s1), cap(s1), &sh1)

	// 对 slice 扩容.
	s1 = append(s1, 4)
	sh1 = (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Printf("s1 扩容后, len: %d, cap: %d, slice header point: %p\n", len(s1), cap(s1), &sh1)
}
