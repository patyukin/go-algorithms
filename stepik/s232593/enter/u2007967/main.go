package main

import "fmt"

const (
	minNumber = 10
	maxNumber = 100
)

func main() {
	for {
		var a int
		fmt.Scan(&a)
		if a < minNumber {
			continue
		}

		if a > maxNumber {
			break
		}

		fmt.Println(a)
	}
}
