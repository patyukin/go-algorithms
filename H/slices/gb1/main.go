package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	printSliceInfo(&slice1, "slice1")
	printSliceInfo(&slice2, "slice2")
}

func printSliceInfo(s *[]int, name string) {
	fmt.Println("---------------------------------------")
	fmt.Printf(
		" | slice: %s,  %v, len: %d, cap: %d, pointer: %v |\n",
		name,
		*s,
		len(*s),
		cap(*s),
		unsafe.SliceData(*s),
	)

	for i := range *s {
		fmt.Printf("pointer of %s %d: %p\n", name, i, &(*s)[i])
	}
}
