package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)

	var nums []int
	for {
		if num < 10 {
			nums = append([]int{num}, nums...)
			break
		}

		a := num % 10
		nums = append([]int{a}, nums...)
		num /= 10
	}

	for k := 0; k < len(nums); k++ {
		fmt.Print(nums[k] * nums[k])
	}
}
