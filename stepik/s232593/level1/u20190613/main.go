package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)

	result := 2
	var findFibonacci func(num int)
	findFibonacci = func(num int) {

		for {
			nextFibonacci := getFibonacci(result + 1)
			if nextFibonacci == a {
				result++
				break
			}

			if nextFibonacci > a {
				result = -1
				break
			}

			result++
		}
	}

	findFibonacci(a)

	fmt.Println(result)
}

func getFibonacci(n int) int {
	if n == 1 {
		return n + 1
	}

	if n == 0 {
		return n + 1
	}

	prevPrev := 0
	prev := 1
	result := 0

	for i := 2; i <= n; i++ {
		result = prevPrev + prev
		prevPrev = prev
		prev = result
	}

	return result
}
