package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Println(getSum(a))
}

func getSum(num int) int {
	a := num % 10
	num /= 10
	b := num % 10
	c := num / 10

	return a + b + c
}
