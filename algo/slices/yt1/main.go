package main

import (
	"fmt"
)

func main() {
	n := make([]int, 1, 2)
	fmt.Println(n) // [1] len = 1, cop = 2

	appendSlice(n, 1024)
	fmt.Println(n) // [1] len = 1, cap = 2

	mutateSlice(n, 0, 512)
	fmt.Println(n) // [512] len = 1, cap = 2
}

func mutateSlice(s []int, idx, val int) {
	s[idx] = val
}

func appendSlice(s []int, val int) {
	s = append(s, val)
}
