package main

import (
	"fmt"
)

func main() {
	var inputString string
	fmt.Scanln(&inputString)

	runes := []rune(inputString)
	left, right := 0, len(runes)-1
	for left < right {
		if string(runes[left]) != string(runes[right]) {
			fmt.Println("Нет")
			return
		}
		left++
		right--
	}

	fmt.Println("Палиндром")
}
