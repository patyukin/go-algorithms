package main

import (
	"fmt"
	"sync"
)

func pipeline() {
	ch1 := produce(2)
	ch3 := produce(10)
	ch := merge(ch1, ch3)
	ch = pow2(ch)
	ch = pow2(ch)
	ch = pow2(ch)
	<-consume(ch)
}

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

func consume(ch <-chan int) chan struct{} {
	res := make(chan struct{})

	go func() {
		for v := range ch {
			fmt.Println(v)
		}

		close(res)
	}()

	return res
}

func splitEven(ch <-chan int) {

}

func fanOut(ch <-chan int) {

}

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	for _, ch := range chs {

		ch := ch
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	fmt.Println("pipeline")
	pipeline()
}
