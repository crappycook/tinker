package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet[string]()
	s.Put("one")
	s.Put("two")
	fmt.Println(s.Has("one"))
	fmt.Println(s.Has("three"))
	s.Delete("one")
	fmt.Println(s.Has("one"))
}
