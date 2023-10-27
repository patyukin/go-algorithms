package main

import (
	"fmt"
)

func main() {
	bits := make(chan int)

	go generateRandomBits(bits, 100000)

	for i := 0; i < 100000; i++ {
		bit := <-bits
		fmt.Print(bit)
	}

}

func generateRandomBits(bits chan<- int, count int) {
	for i := 0; i < count; i++ {
		select {
		case bits <- 0:
		case bits <- 1:
		}
	}
	close(bits)
}
