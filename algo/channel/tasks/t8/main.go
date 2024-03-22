package main

import (
	"fmt"
	"math/rand"
)

func main() {
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			inputChannel <- rand.Intn(5) + 1
		}
		close(inputChannel)
	}()

	go func() {
		for v := range inputChannel {
			outputChannel <- v
		}
		close(outputChannel)
	}()

	var results []int
	for v := range outputChannel {
		results = append(results, v)
	}

	for _, v := range results {
		fmt.Println(v)
	}
}
