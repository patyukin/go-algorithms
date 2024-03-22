package main

import "fmt"

func generator(nums ...int) chan int {
	ch := make(chan int)

	go func() {
		for _, v := range nums {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func consume(ch <-chan int) {
	for v := range ch {
		fmt.Println("v =", v)
	}
}

func main() {
	ch := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	consume(ch)
}
