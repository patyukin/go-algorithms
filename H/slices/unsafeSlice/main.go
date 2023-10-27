package main

// https://go101.org/article/unsafe.html

import (
	"fmt"
	"unsafe"
)

func main() {
	type MyString string
	ms := []MyString{"C", "C++", "Go"}
	fmt.Printf("%s\n", ms) // [C C++ Go]
	// ss := ([]string)(ms) // compiling error
	ss := *(*[]string)(unsafe.Pointer(&ms))
	ss[1] = "Zig"
	fmt.Printf("%s\n", ms) // [C Zig Go]
	// ms = []MyString(ss) // compiling error
	ms = *(*[]MyString)(unsafe.Pointer(&ss))

	// Since Go unsafeSlice.17, we may also use the
	// unsafe.Slice function to do the conversions.
	ss = unsafe.Slice((*string)(&ms[0]), len(ms))
	ss[1] = "С++"
	fmt.Printf("%s\n", ss) // [C С++ Go]
	ms = unsafe.Slice((*MyString)(&ss[0]), len(ms))
	fmt.Printf("%s\n", ms) // [C Zig Go]

	length := len(ms)
	for i := 0; i < length; i++ {
		*(*MyString)(unsafe.Pointer(uintptr(unsafe.Pointer(&ms[0])) + uintptr(i)*unsafe.Sizeof(ms[0]))) = MyString(fmt.Sprintf("%s - %d", "GO", i))
	}

	fmt.Println(ms)
}
