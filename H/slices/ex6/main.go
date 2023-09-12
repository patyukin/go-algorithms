package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	b := s[1:3]

	b[0] = 100
	fmt.Println("s => ", s, "len = ", len(s), "cap = ", cap(s))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b))
	fmt.Println(uintptr(unsafe.Pointer(unsafe.SliceData(s))), uintptr(unsafe.Pointer(unsafe.SliceData(b))))

	b = append(b, 11)
	b[0] = 200
	fmt.Println("s => ", s, "len = ", len(s), "cap = ", cap(s))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b))
	fmt.Println(uintptr(unsafe.Pointer(unsafe.SliceData(s))), uintptr(unsafe.Pointer(unsafe.SliceData(b))))

	b = append(b, 12)
	b[0] = 300
	fmt.Println("s => ", s, "len = ", len(s), "cap = ", cap(s))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b))
	fmt.Println(uintptr(unsafe.Pointer(unsafe.SliceData(s))), uintptr(unsafe.Pointer(unsafe.SliceData(b))))

	b = append(b, 13)
	b[0] = 400
	fmt.Println("s => ", s, "len = ", len(s), "cap = ", cap(s))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b))
	fmt.Println(uintptr(unsafe.Pointer(unsafe.SliceData(s))), uintptr(unsafe.Pointer(unsafe.SliceData(b))))
}
