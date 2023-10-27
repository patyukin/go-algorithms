package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scanln(&str)

	runes := []rune(str)
	var newRunes []rune
	for i, v := range runes {
		newRunes = append(newRunes, v)
		if i == len(runes)-1 {
			continue
		}

		newRunes = append(newRunes, '*')
	}

	fmt.Println(string(newRunes))
}
