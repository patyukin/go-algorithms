package main

import "fmt"

// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	first, second := 0, 1

	for i := 0; i < n; i++ {
		first, second = second, first+second
	}

	return first
}

func main() {
	fmt.Println(fib(3))
	fmt.Println(fib2(3))

}
