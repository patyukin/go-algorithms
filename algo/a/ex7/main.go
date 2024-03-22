package main

import (
	"fmt"
	"strconv"
)

func removeDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	var i int
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return nums[:i+1]
}

func compressedUnicode(s string) string {
	var prev rune
	var result string
	count := 1

	for idx, current := range s {
		if idx == 0 {
			prev = current
			continue
		}

		if prev == current {
			count++
			continue
		}

		result += fmt.Sprintf("%c%d", prev, count)
		count = 1
		prev = current
	}

	result += fmt.Sprintf("%c%d", prev, count)

	return result
}

func compressed(s string) string {
	var result string
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			count++
		} else {
			result += string(s[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	result += string(s[len(s)-1]) + strconv.Itoa(count)

	if len(result) < len(s) {
		return result
	}

	return s
}

func compressString(s string) string {
	var result string

	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++
		}

		result += string(s[i]) + strconv.Itoa(count)
	}

	return result
}

func compress2(s string) string {
	var result string
	var count int

	for i := 0; i < len(s); i++ {
		count++

		if i+1 == len(s) || s[i] != s[i+1] {
			result += string(s[i]) + strconv.Itoa(count)
			count = 0
		}
	}

	if len(result) < len(s) {
		return result
	}

	return s
}

func compressModify(s string) string {
	count := 1
	var result string
	var prev rune

	for idx, current := range s {
		if idx == 0 {
			prev = current
			continue
		}

		if prev == current {
			count++
			continue
		}

		if count > 1 {
			result += fmt.Sprintf("%c%d", prev, count)
		} else {
			result += string(prev)
		}
		count = 1
		prev = current
	}

	if count > 1 {
		result += fmt.Sprintf("%c%d", prev, count)
	} else {
		result += string(prev)
	}

	return result
}
