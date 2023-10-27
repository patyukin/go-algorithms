package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 0, 3)

	s = append(s, 1, 2)
	fmt.Println("s => ", s, "len = ", len(s), "cap = ", cap(s))
	fmt.Println(uintptr(unsafe.Pointer(unsafe.SliceData(s))), "\n")

	s = s[:3]
	fmt.Println("s => ", s, "len = ", len(s), "cap = ", cap(s))
	fmt.Println(uintptr(unsafe.Pointer(unsafe.SliceData(s))), "\n")

	foo(s)
	fmt.Println("s => ", s, "len = ", len(s), "cap = ", cap(s))
	fmt.Println(uintptr(unsafe.Pointer(unsafe.SliceData(s))), "\n")
}

func foo(s []int) {
	s = append(s, 3)

	fmt.Println("s => ", s, "len = ", len(s), "cap = ", cap(s))
	fmt.Println(uintptr(unsafe.Pointer(unsafe.SliceData(s))), "\n")
}
