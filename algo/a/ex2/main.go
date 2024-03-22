package main

import (
	"math/rand"
	"time"
)

func main() {
	for num := range generator(10) {
		println(num)
	}
}

func generator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- r.Intn(n)
		}

		close(ch)
	}()

	return ch
}
