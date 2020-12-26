package queue

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := New(6)
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	for key, count := range items {
		heap.Push(pq, &Item{
			Value:    key,
			Priority: int64(count),
		})
	}
	for i := 0; i < 3; i++ {
		item := heap.Pop(pq).(*Item)
		fmt.Println("item:", item.Value, item.Priority)
	}
}
