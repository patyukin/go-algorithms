package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	roots, status := solveEquation(a, b, c)
	if status != "" {
		fmt.Println(status)
	} else {
		for _, r := range roots {
			fmt.Println(r)
		}
	}
}

func solveEquation(a, b, c int) ([]int, string) {
	if c < 0 {
		return nil, "NO SOLUTIONS"
	}

	if a == 0 && b == 0 && c == 0 {
		return nil, "MANY SOLUTIONS"
	}

	if a != 0 && b == 0 && c != 0 {
		r := c * c / a
		return []int{-r, r}, "NO SOLUTIONS"
	}

	return []int{(c*c - b) / a}, ""
}
