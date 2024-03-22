package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := make([]int, 0, 3)
	slice = append(slice, 1<<3)

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(sliceHeader))
	fmt.Println("-----------------------------------------")
	fmt.Printf(
		"| slice: %v, size: %d, sliceHeader: %v |\n",
		slice,
		unsafe.Sizeof(slice),
		sliceHeader,
	)

	func(inLambdaSlice []int) {
		lambdaHeader := (*reflect.SliceHeader)(unsafe.Pointer(&inLambdaSlice))
		fmt.Printf(
			"| in lambda: sliceHeader was copied BY VALUE: %v|\n",
			lambdaHeader,
		)

		inLambdaSlice[0] = 1 << 4
		inLambdaSlice = append(inLambdaSlice, 1<<5)
		inLambdaSlice = append(inLambdaSlice, 1<<6)

		fmt.Printf("| in lambda: sliceHeader after appending: %v |\n", lambdaHeader)
		fmt.Printf("| in lambda: slice: %v |\n", inLambdaSlice)
		fmt.Println("-----------------------------------------")
	}(slice)

	slice = slice[:cap(slice)]
	fmt.Printf("| in lambda: slice: %v |\n", slice)
}
