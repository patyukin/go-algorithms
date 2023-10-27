package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type MyString string
	ms := []MyString{"C", "C++", "Go"}

	nsd := unsafe.SliceData(ms)
	ptr := uintptr(unsafe.Pointer(nsd))
	ptr += unsafe.Sizeof(*nsd)
	ptr2 := ptr
	fmt.Println(*(*MyString)(unsafe.Pointer(ptr2)))
}
