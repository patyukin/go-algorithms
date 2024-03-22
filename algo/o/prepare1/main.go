package main

import "sort"

// https://leetcode.com/problems/maximum-product-of-three-numbers/

func maxNumberProductWithSort(s []int) int {
	var temp []int
	for i := 1; i < len(s); i++ {
		temp = append(temp, s[i-1]*s[i])
	}

	sort.Ints(temp)
	return temp[len(temp)-1]
}

func maxNumberProductWithoutSort(s []int) int {
	result := s[0] * s[1]
	for i := 2; i < len(s); i++ {
		if result < s[i]*s[i-1] {
			result = s[i] * s[i-1]
		}
	}

	return result
}
