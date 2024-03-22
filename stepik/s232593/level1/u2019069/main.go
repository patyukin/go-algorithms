package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)
	sumNum := getSumNumber(num)
	fmt.Println(getSumNumber(sumNum))
}

func getSumNumber(num int) int {
	var sum int
	for {
		if num < 10 {
			sum += num
			break
		}

		a := num % 10
		sum += a
		num /= 10
	}

	return sum
}
