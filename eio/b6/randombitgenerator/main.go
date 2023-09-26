package main

import (
	"fmt"
)

func generateBit(ch chan int) {
	for i := 0; i < 100000; i++ {
		switch i%2 == 0 {
		case true:
			ch <- 0
		default:
			ch <- 1
		}
	}

	close(ch)
}

func main() {
	ch := make(chan int)
	go generateBit(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
