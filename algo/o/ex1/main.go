package main

import (
	"fmt"
	"sort"
)

// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate

func containsDuplicates(nums []int) bool {
	m := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = struct{}{}
			continue
		}

		return true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicates([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicates([]int{1, 2, 3}))
}

func containsDuplicates2(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
