package main

import "fmt"

func main() {
	var num, a, b int
	fmt.Scanln(&num)

	a = num / 3600
	num -= a * 3600
	b = num / 60

	fmt.Printf("It is %d hours %d minutes.", a, b)
}
