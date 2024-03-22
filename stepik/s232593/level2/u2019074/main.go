package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scanln(&str)

	runes := []rune(str)
	maxNumber := string(runes[0])
	for i := 1; i < len(runes); i++ {
		if string(runes[i]) > maxNumber {
			maxNumber = string(runes[i])
		}
	}

	fmt.Println(maxNumber)
}
