package alg

import (
	"fmt"
	"testing"
)

func TestSort(T *testing.T) {
	fmt.Println(sort([]int{1, 3, 6, 7, 8}, []int{2, 5, 9}))
}

func TestLruCache(t *testing.T) {
	cache := NewLruCache(5)
	_ = cache.Set("k1", "v1")
	_ = cache.Set("k2", "v2")
	_ = cache.Set("k3", "v3")
	_ = cache.Set("k4", "v4")
	_ = cache.Set("k5", "v5")
	err := cache.Set("k6", "v6")
	fmt.Println(err)
}
