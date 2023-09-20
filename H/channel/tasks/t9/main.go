package main

import "fmt"

func main() {
	dataChannel := make(chan int)
	resultChannel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			dataChannel <- i
		}

		close(dataChannel)
	}()

	go func() {
		for v := range dataChannel {
			resultChannel <- v * v
		}

		close(resultChannel)
	}()

	result := make([]int, 0)
	for v := range resultChannel {
		result = append(result, v)
	}

	fmt.Println(result)
}
