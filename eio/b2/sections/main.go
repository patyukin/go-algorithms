package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
LOOP:
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto LOOP
		}
		fmt.Println(i)
	}
}
