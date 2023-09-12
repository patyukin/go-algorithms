package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	if a*a+b*b == c*c || a*a+c*c == b*b || c*c+b*b == a*a {
		fmt.Println("Прямоугольный")
		return
	}

	fmt.Println("Непрямоугольный")
}
