package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Scanln(&password)

	runes := []rune(password)
	if len(runes) < 5 {
		fmt.Println("Wrong password")
		return
	}

	for _, r := range runes {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			fmt.Println("Wrong password")
			return
		}
	}

	fmt.Println("Ok")
}
