package lru

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := New(2)
	lru.Put("one", 1)
	lru.Put("two", 2)
	lru.Put("three", 3)
	fmt.Println(lru.Get("one"))
	fmt.Println(lru.Get("two"))
	fmt.Println(lru.Len())
}
