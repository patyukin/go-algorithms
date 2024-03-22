package main

// 136. Single Number
// https://leetcode.com/problems/single-number

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}
