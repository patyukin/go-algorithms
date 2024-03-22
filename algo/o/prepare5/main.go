package main

// Input: a="лапоть", b="пальто"
// Output: true
// https://leetcode.com/problems/valid-anagram/
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, value := range s {
		m[value]++
	}

	for _, value := range t {
		if m[value] == 0 {
			return false
		}

		m[value]--
	}

	return true
}
