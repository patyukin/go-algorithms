package main

import "fmt"

// 896. Monotonic Array
func isMonotonic(nums []int) bool {
	increasing := true
	decreasing := true

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			decreasing = false
		}
		if nums[i] < nums[i-1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
}
