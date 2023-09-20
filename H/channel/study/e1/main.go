package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	fmt.Println(ch)

	ch <- 100
	ch <- 100
	ch <- 100
	ch <- 100

	fmt.Println(<-ch)
}
