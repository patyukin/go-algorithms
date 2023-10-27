package main

import (
	"fmt"
	"sync"
)

func produce(x int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := x; i < x+2; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func pow2(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func fanOutPow2(in chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func consume(chs ...<-chan int) chan struct{} {
	res := make(chan struct{})
	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)
		ch := ch
		go func() {
			defer wg.Done()
			for v := range ch {
				fmt.Println(v)
			}

		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
