package main

import (
	"fmt"
)

func main() {
	arr := [...]int{9, 7, 5, 3, 1}
	nums := arr[2:]
	nums2 := nums[1:]

	arr[2]++
	nums[1] -= arr[4] - 4
	nums2[1] += 5

	fmt.Println(nums)
}
