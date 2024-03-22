package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/top-k-frequent-elements
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var n [][2]int
	for i, count := range m {
		n = append(n, [2]int{i, count})
	}

	sort.Slice(n, func(i, j int) bool {
		return n[i][1] > n[j][1]
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = n[i][0]
	}

	return result
}

func topKFrequent2(nums []int, k int) []int {
	var res []int
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	temp := make([][]int, len(nums)+1)
	for key, count := range m {
		temp[count] = append(temp[count], key)
	}

	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] != nil {
			res = append(res, temp[i]...)
			if len(res) >= k {
				res = res[:k]
				break
			}
		}
	}

	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 3, 3, 3, 3, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent2([]int{1, 3, 3, 3, 3, 1, 1, 2, 2, 3}, 2))
}
