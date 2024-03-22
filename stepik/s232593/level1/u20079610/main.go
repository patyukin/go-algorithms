package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber, secondNumber int
	fmt.Scanln(&firstNumber, &secondNumber)
	first := getNumbers(firstNumber)
	second := getNumbers(secondNumber)
	var result string

	for _, f := range first {
		if contains(second, f) {
			prefix := " "
			if result == "" {
				prefix = ""
			}

			result += prefix + strconv.Itoa(f)
		}
	}

	fmt.Println(result)
}

func getNumbers(num int) []int {
	var temp []int

	for {
		if num < 10 {
			temp = append([]int{num}, temp...)
			break
		}

		a := num % 10
		temp = append([]int{a}, temp...)
		num /= 10
	}

	return temp
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
