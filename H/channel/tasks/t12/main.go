package main

import (
	"fmt"
)

func calculateFactorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	results := make(chan int)

	go func() {
		results <- calculateFactorial(5)
	}()
	go func() {
		results <- calculateFactorial(7)
	}()

	fmt.Println(<-results)
	fmt.Println(<-results)
}
