package main

import "fmt"

func main() {
	var num, a, b, c int
	fmt.Scanln(&num)

	a = num % 10
	num /= 10
	b = num % 10
	num /= 10
	c = num

	fmt.Printf("%d%d%d", a, b, c)
}
