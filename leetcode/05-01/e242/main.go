package main

import "fmt"

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	for _, char := range t {
		m[char]++
	}

	for _, char := range s {
		if m[char] == 0 {
			return false
		}

		m[char]--
	}

	return true
}
