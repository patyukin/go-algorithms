package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:]

	b[0] = 100
	fmt.Println("a => ", a, "len = ", len(a), "cap = ", cap(a), "uintptr = ",
		uintptr(unsafe.Pointer(&a)))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b), "uintptr = ",
		uintptr(unsafe.Pointer(unsafe.SliceData(b))))

	b = append(b, 11)
	b[0] = 200
	fmt.Println("a => ", a, "len = ", len(a), "cap = ", cap(a), "uintptr = ",
		uintptr(unsafe.Pointer(&a)))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b), "uintptr = ",
		uintptr(unsafe.Pointer(unsafe.SliceData(b))))

	b = append(b, 12)
	b[0] = 300
	fmt.Println("a => ", a, "len = ", len(a), "cap = ", cap(a), "uintptr = ",
		uintptr(unsafe.Pointer(&a)))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b), "uintptr = ",
		uintptr(unsafe.Pointer(unsafe.SliceData(b))))

	b = append(b, 13)
	b[0] = 400
	fmt.Println("a => ", a, "len = ", len(a), "cap = ", cap(a), "uintptr = ",
		uintptr(unsafe.Pointer(&a)))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b), "uintptr = ",
		uintptr(unsafe.Pointer(unsafe.SliceData(b))))
}
