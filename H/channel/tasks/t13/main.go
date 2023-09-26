package main

import (
	"fmt"
)

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			first <- i
		}
		close(first)
	}()

	go func() {
		for i := 1; i < 10; i++ {
			second <- i
		}
		close(second)
	}()

	go func() {
		for i := 1; i < 10; i++ {
			third <- i
		}
		close(third)
	}()

	for {
		select {
		case v, ok := <-first:
			if !ok {
				return
			}
			fmt.Println(v)
		case v, ok := <-second:
			if !ok {
				return
			}
			fmt.Println(v)
		case v, ok := <-third:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}
}
