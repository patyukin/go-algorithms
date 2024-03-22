package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 0, 3)

	s = append(s, 1, 2)
	nsd := unsafe.SliceData(s)
	ptr := uintptr(unsafe.Pointer(nsd))
	fmt.Println("before foo", ptr) // before foo 824633844264
	fmt.Println("before foo", s)   // 1, 2, 3 len = 3, cap = 3

	foo(&s)

	nsd = unsafe.SliceData(s)
	ptr = uintptr(unsafe.Pointer(nsd))
	fmt.Println("after foo", ptr) // after foo 824633844264
	fmt.Println("after foo", s)   // 1, 2, 3 len = 3, cap = 3
}

func foo(s *[]int) {
	*s = append(*s, 3)

	nsd := unsafe.SliceData(*s)
	ptr := uintptr(unsafe.Pointer(nsd))
	fmt.Println("foo", ptr) // foo 824633844264
	fmt.Println("foo", *s)  // 1, 2, 3 len = 3, cap = 3
}
