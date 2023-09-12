package main

import (
	"fmt"
	"reflect"
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

	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	header.Cap = 100
	fmt.Println("a => ", a, "len = ", len(a), "cap = ", cap(a), "uintptr = ",
		uintptr(unsafe.Pointer(&a)))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b), "uintptr = ",
		uintptr(unsafe.Pointer(unsafe.SliceData(b))))

	b = append(b, 10, 11, 12, 13, 14, 15, 16)
	fmt.Println("a => ", a, "len = ", len(a), "cap = ", cap(a), "uintptr = ",
		uintptr(unsafe.Pointer(&a)))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b), "uintptr = ",
		uintptr(unsafe.Pointer(unsafe.SliceData(b))))

	for range b {

	}

	b = append(b, 10, 11, 12, 13, 14, 15, 16)
	fmt.Println("a => ", a, "len = ", len(a), "cap = ", cap(a), "uintptr = ",
		uintptr(unsafe.Pointer(&a)))
	fmt.Println("b => ", b, "len = ", len(b), "cap = ", cap(b), "uintptr = ",
		uintptr(unsafe.Pointer(unsafe.SliceData(b))))
}
