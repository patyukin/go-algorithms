package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	intSlice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		intSlice = append(intSlice, i)
	}

	printSliceInfo(&intSlice, "intSlice")

	cut := intSlice[2:4]
	printSliceInfo(&cut, "cut")

	cut = cut[:cap(cut)]
	_ = cut[2]
	printSliceInfo(&cut, "cut")

	fmt.Printf(
		"| intSlice: %d, cut: %d |\n"+
			"| intSlice: %v, cut: %v |\n",
		reflect.ValueOf(intSlice).Pointer(),
		reflect.ValueOf(cut).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&intSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cut)),
	)

	cut = append(cut, 1<<32+1)
	printSliceInfo(&cut, "cut")
}

func printSliceInfo(s *[]int, name string) {
	fmt.Println("---------------------------------------")
	fmt.Printf(
		" | intSlice slice %v, len: %d, cap: %d, pointer: %v |\n",
		*s,
		len(*s),
		cap(*s),
		unsafe.SliceData(*s),
	)

	for i := range *s {
		fmt.Printf("pointer of %s %d: %p, %d\n", name, i, &(*s)[i], pointerToDec(&(*s)[i]))
	}
}

func pointerToDec(ptr *int) uint64 {
	ptrInt := uintptr(unsafe.Pointer(ptr))
	ptrDec := uint64(ptrInt)

	return ptrDec
}
