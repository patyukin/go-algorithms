package main

import (
	"fmt"
	"strconv"
)

// 7. Reverse Integer
// https://leetcode.com/problems/reverse-integer
func reverse(x int) int {
	result := 0
	for ; x != 0; x /= 10 {
		result = result*10 + x%10
	}

	if result > 1<<31-1 || result < -(1<<31) {
		return 0
	}

	return result
}

func main() {

	fmt.Println(reverse(123))
	fmt.Println(reverse2(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse2(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse2(120))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse2(1534236469))

}

func power(base, exponent int) int {
	result := 1

	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func reverse2(x int) int {
	if x == 0 {
		return 0
	}

	s := 1
	if x < 0 {
		s = -1
		x = -x
	}

	var result int
	var preResult string

	for {
		if x < 10 {
			preResult = preResult + fmt.Sprintf("%d", x)
			break
		}

		temp := x % 10
		preResult = preResult + fmt.Sprintf("%d", temp)
		x = x / 10
	}

	result, _ = strconv.Atoi(preResult)
	result *= s

	if result > 1<<31-1 || result < -(1<<31) {
		return 0
	}

	return result
}
