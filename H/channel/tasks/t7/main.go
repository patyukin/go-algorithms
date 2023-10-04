package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	producer := make(chan int)
	go func() {
		for v := range producer {
			fmt.Println(v)
		}
		close(producer)
	}()

	go func() {
		for {
			producer <- rand.Intn(5) + 1
			time.Sleep(time.Second)
		}
	}()

}
