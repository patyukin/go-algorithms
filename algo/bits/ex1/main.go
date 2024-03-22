package main

import "fmt"

func getBinary(n int) int {
	var result int
	for n > 0 {
		temp := n % 2
		if temp == 1 {
			result++
		}

		n /= 2
	}

	return result
}

func main() {
	fmt.Println(getBinary(5))
	fmt.Println(getBinary(11))
	fmt.Println(getBinary(0))
}
