package main

import (
	"errors"
	"fmt"
)

type ArrQueue struct {
	len       int
	headIndex int // 头下标，出队列.
	tailIndex int // 尾下班，入队列.
	arr       []int
}

func initArrQueue(len int) *ArrQueue {
	return &ArrQueue{
		len:       len,
		headIndex: 0,
		tailIndex: 0,
		arr:       make([]int, len),
	}
}

func (q *ArrQueue) enqueue(data int) error {
	if q.tailIndex == q.len {
		return errors.New("queue is full")
	}

	q.arr[q.tailIndex] = data
	q.tailIndex++
	return nil
}

func (q *ArrQueue) dequeue() (int, error) {
	if q.headIndex == q.tailIndex {
		return 0, errors.New("queue is empty")
	}

	data := q.arr[q.headIndex]
	q.headIndex++
	return data, nil
}

func main() {
	que := initArrQueue(5)
	_ = que.enqueue(1)
	_ = que.enqueue(2)
	_ = que.enqueue(3)
	_ = que.enqueue(4)
	err := que.enqueue(5)
	fmt.Println(err)
	err = que.enqueue(6)
	fmt.Println(err)
	for i := 0; i <= 5; i++ {
		fmt.Println(que.dequeue())
	}
}
