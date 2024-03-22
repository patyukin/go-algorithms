package main

// Дан массив чисел. Нужно найти в нём последовательный под-массив, который имеет наибольшую сумму элементов. Вернуть эту сумму.
// Пример

// leetcode.com/problems/maximum-subarray

// Input: nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// Output: 6
// Explanation: sub-array = [4, -1, 2, 1],  sum = 6

func maxSubArray(nums []int) int {
	currentSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum += nums[i]
		if nums[i] > currentSum {
			currentSum = nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func getSum2(nums []int) int {
	currentSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(currentSum, maxSum)
	}

	return maxSum
}

func getSum(nums []int) int {
	s := nums[0]

	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		for j := i + 1; j < len(nums); j++ {
			temp += nums[j]
		}

		if temp > s {
			s = temp
		}
	}

	return s
}
