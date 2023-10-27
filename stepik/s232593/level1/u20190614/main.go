package main

import "fmt"

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		remainder := n % 2
		binary = fmt.Sprint(remainder) + binary
		n = n / 2
	}

	return binary
}

func main() {
	var a int
	fmt.Scanln(&a)

	binary := toBinary(a)
	fmt.Println(binary)
}
