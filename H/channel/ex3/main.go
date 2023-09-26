package main

import (
	"fmt"
)

func generateNumber(min, max int, out chan int) chan int {
	for i := min; i <= max; i++ {
		out <- i
	}
	close(out)

	return out
}

func main() {
	minNumber, maxNumber := 2, 4
	out := make(chan int)

	go generateNumber(minNumber, maxNumber, out)

	sum := 0
	for v := range out {
		sum += v * v
	}

	fmt.Println(sum)
}
