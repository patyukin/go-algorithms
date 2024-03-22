package main

import (
	"fmt"
	"sort"
)

// 977. Squares of a Sorted Array
// https://leetcode.com/problems/squares-of-a-sorted-array

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares2([]int{-4, -1, 0, 3, 10}))

	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
	fmt.Println(sortedSquares2([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares2(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func sortedSquares(nums []int) []int {
	var result []int
	for _, v := range nums {
		result = append(result, v*v)
	}

	sort.Ints(result)
	return result
}
