package main

import "fmt"

// Даны две строки, содержащие буквы английского алфавита и символ #,
// соответствующий клавише backspace. Нужно сравнить две строки с учётом backspace и вернуть TRUE если они равны.
// Пример

// Input: s = "ab#c", t = "ad#c"
// Output: true
// Explanation: Both s and t become "ac"

// leetcode.com/problems/backspace-string-compare

func backspaceCompare(s, t string) bool {
	var getString func(string) string
	getString = func(s string) string {
		var res []rune
		for _, c := range s {
			if c == '#' {
				if len(res) > 0 {
					res = res[:len(res)-1]
				}
			} else {
				res = append(res, c)
			}
		}

		return string(res)
	}

	return getString(s) == getString(t)
}

func main() {
	fmt.Println(check("ab##", "c#d###"))
}

func check(s, t string) bool {
	var getString func(string) string
	getString = func(ss string) string {
		var newS string
		for i := 1; i < len(s); i++ {
			if s[i] == '#' {
				if i-1 < 0 {
					continue
				}

				newS += newS[:len(newS)-1]
			} else {
				newS += string(s[i])
			}
		}

		return newS
	}

	return getString(s) == getString(t)
}
