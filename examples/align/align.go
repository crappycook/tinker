package main

import (
	"fmt"
	"unsafe"
)

type t1 struct {
	a int8
	b int16
	c int32
}

// 字段排序不同 所占内存大小不同
type t2 struct {
	a int8
	c int32
	b int16
}

// 空 struct{} 内存对齐
type x struct {
	a int
	b struct{}
}

type y struct {
	a struct{}
	b int
}

func main() {
	m, n := t1{}, t2{}
	fmt.Println(unsafe.Alignof(m)) // 4
	fmt.Println(unsafe.Alignof(n)) // 4
	fmt.Println(unsafe.Sizeof(m))  // 8
	fmt.Println(unsafe.Sizeof(n))  // 12

	i, j := x{}, y{}
	fmt.Println(unsafe.Alignof(i)) // 8
	fmt.Println(unsafe.Alignof(j)) // 8
	fmt.Println(unsafe.Sizeof(i))  // 16
	fmt.Println(unsafe.Sizeof(j))  // 8
}
