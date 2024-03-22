package main

import (
	"fmt"
	"strconv"
)

func compressString(s string) string {
	result := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += string(s[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	// Добавление последнего символа и его счетчика
	result += string(s[len(s)-1]) + strconv.Itoa(count)

	// Если сжатая строка короче исходной, вернуть сжатую строку,
	// иначе вернуть исходную строку
	if len(result) < len(s) {
		return result
	} else {
		return s
	}
}

func main() {
	s := "aaabbbcaa"
	compressed := compressString(s)
	fmt.Println(compressed) // Вывод: "a3b3c3"
}
