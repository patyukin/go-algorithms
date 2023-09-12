package main

import (
	"fmt"
	"unsafe"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	ptr := &numbers[0]
	length := len(numbers)
	capacity := cap(numbers)

	array := (*[5]int)(unsafe.Pointer(ptr))

	fmt.Println("Array:", *array)
	fmt.Println("Length:", length)
	fmt.Println("Capacity:", capacity)
}
