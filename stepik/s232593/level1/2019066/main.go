package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	if a+b > c && a+c > b && c+b > a {
		fmt.Println("Существует")
		return
	}

	fmt.Println("Не существует")
}
