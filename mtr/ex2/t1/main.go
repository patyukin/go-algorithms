package main

import (
	"fmt"
	"unsafe"
)

type account struct {
	value int
}

func main() { //
	s1 := make([]account, 0, 2) // len = 0, cap = 2, s = []

	nsd := unsafe.SliceData(s1)
	ptr := uintptr(unsafe.Pointer(nsd))
	fmt.Println("before append, ptr = ", ptr) // before append 1374390648864
	fmt.Println("before append, s1 = ", s1)   // [], len = 0, cap = 2

	s1 = append(s1, account{}) // len = 1, cap = 2, s1 = [account{}]

	nsd2 := unsafe.SliceData(s1)
	ptr2 := uintptr(unsafe.Pointer(nsd2))
	fmt.Println("before append, ptr2 = ", ptr2)
	fmt.Println("before append, s1 = ", s1)

	s2 := append(s1, account{})

	nsd3 := unsafe.SliceData(s1)
	ptr3 := uintptr(unsafe.Pointer(nsd3))
	fmt.Println("before append, ptr3 = ", ptr3)
	fmt.Println("before append, s2 = ", s2)

	// s1: len = 1, cap = 2, ptr = [account{}, account{}] => s1[1] = account{}
	// s2: len = 2, cap = 2, ptr = [account{}, account{}]

	// s1.ptr == s2.ptr - true

	acc := &s2[0]
	// acc = &account{}

	acc.value = 100
	// acc = &account{value: 100} (s2)

	fmt.Println(s1, s2)
	// s1: len = 1, cap = 2, ptr = [account{value: 100}]
	// s2: len = 2, cap = 2, ptr = [account{value: 100}, account{}]

	// [{100}] [{100}, {0}]

	s1 = append(s2, account{})
	// s1: len = 3, cap = 4, ptr = [account{value: 100}, account{}, account{}]
	// s2: len = 2, cap = 2, ptr = [account{value: 100}, account{}]

	acc.value += 100
	// s1: len = 3, cap = 4, ptr = [account{value: 100}, account{}, account{}]
	// s2: len = 2, cap = 2, ptr = [account{value: 200}, account{}]

	fmt.Println(s1, s2)

	// [{100}, {}, {}], [{200}, {}]
}
