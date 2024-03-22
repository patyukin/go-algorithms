package main

import "fmt"

func main() {
	switch fmt.Print("a"); {
	default:
		fmt.Print("b")
	case true:
		fmt.Print("c")
	}
}
