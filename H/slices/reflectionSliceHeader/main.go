package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := []int{1, 2, 3}

	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Println("Length:", header.Len)            // Длина
	fmt.Println("Capacity:", header.Cap)          // Вместимость
	fmt.Println("Pointer to array:", header.Data) // Указатель на массив

	for i := 0; i < int(header.Len); i++ {
		value := *(*int)(unsafe.Pointer(header.Data + uintptr(i)*unsafe.Sizeof(slice[0])))
		fmt.Println("Value at index", i, "is", value)
		*(*int)(unsafe.Pointer(header.Data + uintptr(i)*unsafe.Sizeof(slice[0]))) = 10 + i
	}

	fmt.Println("Pointer to array:", header.Data) // Указатель на массив
	fmt.Println(*(*int)(unsafe.Pointer(header.Data)))

	ptr := uintptr(unsafe.Pointer(&slice[0]))
	for i := 0; i < len(slice); i++ {
		value := *(*int)(unsafe.Pointer(ptr))
		fmt.Println(value)
		ptr += unsafe.Sizeof(slice[0])
	}
}
