package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(`NO`)
}
