package main

// 125. Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

func isLetterOrDigit(char byte) bool {
	return (char >= 65 && char <= 90) || (char >= 97 && char <= 122) || (char >= 48 && char <= 57)
}

func toLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}

	return char
}

func isPalindrome1(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isLetterOrDigit(s[left]) {
			left++
			continue
		}

		if !isLetterOrDigit(s[right]) {
			right--
			continue
		}

		if toLower(rune(s[left])) == toLower(rune(s[right])) {
			left++
			right--
			continue
		}

		return false
	}

	return true
}

func isPalindrome(s string) bool {
	var s2 string
	for i := 0; i < len(s); i++ {
		if !isLetterOrDigit(s[i]) {
			continue
		}

		if s[i] >= 'A' && s[i] <= 'Z' {
			s2 += string(s[i] - 'A' + 'a')
			continue
		}

		s2 += string(s[i])
	}

	if s2 == "" {
		return true
	}

	left, right := 0, len(s2)-1
	for left < right {
		if s2[left] != s2[right] {
			return false
		}

		left++
		right--
	}

	return true
}
