package main

import (
	"fmt"
	"unsafe"
)

type account struct {
	value int
}

// func append([]T, ....T) []T

// len
// cap
// unsafe.Pointer

type Slice struct {
	len int
	cap int
	ptr unsafe.Pointer
}

// s1.ptr == s2.ptr
// s1.len != s2.len

// append(s1, t) s2
// s1 {1, 2, ptr}
// s1 := append(s1, t)
// s1 {2, 2, ptr}
// НО
// s1 {1, 2, ptr}
// s2 := append(s1, t)
// s1 {1, 2, ptr}
// s2 {2, 2, ptr}

// s.len == 2
func foo(s *[]int) {
	(*s)[1] = 100500
	*s = append(*s, 1)
	*s = append(*s, 1, 3, 4, 3, 5)
}

func main() {
	s := make([]int, 0, 3)
	s = append(s, 1)
	s = append(s, 2)

	foo(&s)
	fmt.Println(s)
}

// s.len == 2
//func append(s Slice, t ...Slice) Slice {
//	if s.len != s.cap {
//		s.ptr[s.len] = t
//	}
//	return Slice{
//		len: s.len + 1,
//		cap: s.cap,
//		ptr: s.ptr,
//	}
//}
//
//func append2(s *Slice, t T) {
//	if s.len != s.cap {
//		s.ptr[s.len] = t
//		s.len++
//	}
//}
