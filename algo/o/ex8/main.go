package main

import (
	"fmt"
)

func merge(arrs ...[]int) []int {
	var res []int
	for _, arr := range arrs {
		res = append(res, arr...)
	}

	return res
}

func merge3(arrs ...[]int) []int {
	var length int
	for _, arr := range arrs {
		length += len(arr)
	}

	res := make([]int, length)
	h := res[:]
	for _, arr := range arrs {
		copy(h, arr)
		h = h[len(arr):]
	}

	return res
}

func merge2(arrs ...[]int) []int {
	var length int
	for _, arr := range arrs {
		length += len(arr)
	}

	res := make([]int, length)
	for i, arr := range arrs {
		for _, v := range arr {
			res[i] = v
		}
	}

	return res
}

func main() {
	fmt.Println(merge3([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}))
}
