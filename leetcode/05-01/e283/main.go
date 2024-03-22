package main

import "fmt"

// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/
func moveZeroes(nums []int) []int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	return nums
}

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
}
