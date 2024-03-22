package main

import "sort"

// Input: nums = [0, 1, 1, 6, 12, 20]
// Output: squares = [0, 1, 1, 36, 144, 400]

// https://leetcode.com/problems/squares-of-a-sorted-array/
func getPow(n []int) []int {
	var result []int

	for i := 0; i < len(n); i++ {
		result = append(result, n[i]*n[i])
	}

	sort.Ints(result)

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func getPow2(nums []int) []int {
	result := make([]int, len(nums))
	n := len(nums)
	left, right := 0, n-1
	for i := n - 1; i >= 0; i-- {
		if abs(nums[left]) < abs(nums[right]) {
			result[i] = nums[right] * nums[right]
			right--
		} else {
			result[i] = nums[left] * nums[left]
			left++
		}
	}

	return result
}
