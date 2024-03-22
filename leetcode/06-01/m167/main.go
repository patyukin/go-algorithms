package main

// 167. Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		left := i + 1
		right := len(numbers) - 1

		for left <= right {
			middle := (left + right) / 2
			if numbers[middle] == target-numbers[i] {
				return []int{i + 1, middle + 1}
			} else if numbers[middle] < target-numbers[i] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return []int{}
}
