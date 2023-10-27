package main

import (
	"fmt"
)

func main() {
	fmt.Println(min("long", "Short"))
	for k, v := range []rune("long") {
		fmt.Println(k, v)
	}

	fmt.Println("")

	for k, v := range []rune("Short") {
		fmt.Println(k, v)
	}
}
