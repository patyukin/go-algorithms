package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)
	lastNum := getLastNumber(a)

	if lastNum == 1 && a != 11 {
		fmt.Printf("%d korova", a)
		return
	}

	if lastNum == 2 && a != 12 || lastNum == 3 && a != 13 || lastNum == 4 && a != 14 {
		fmt.Printf("%d korovy", a)
		return
	}

	fmt.Printf("%d korov", a)
}

func getLastNumber(num int) int {
	if num < 10 {
		return num
	}

	return num % 10
}
