package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 3)
	s = append(s, 1, 2)
	fmt.Println("before foo", s) // [1, 2] len = 2, cap = 3

	foo(s)
	fmt.Println("after foo", s) // [100, 200] len = 2, cap = 3
}

func foo(s []int) {
	s[0] = 100
	s[1] = 200
	fmt.Println("foo", s) // [100, 200] len = 2, cap = 3
}
