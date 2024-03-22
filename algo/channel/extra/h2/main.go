package main

import (
	"fmt"
	"time"
)

func splitFull() {
	fmt.Println("split full")
	ch1 := produce(2)
	ch3 := produce(10)
	ch := merge(ch1, ch3)
	even, odd := split(ch)

	even = pow2(even)
	odd = pow2(odd)
	<-consume(odd, even)
}

func fanOutFull() {
	chs := fanOut(make(chan int), 3)

	in := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			in <- i
		}
		close(in)
	}()

	for _, ch := range chs {
		go func(ch <-chan int) {
			for v := range ch {
				fmt.Println("Received:", v)
			}
		}(ch)
	}

	go func() {
		for v := range in {
			for _, ch := range chs {
				ch <- v
			}
		}

		// Закрыть все созданные каналы
		for _, ch := range chs {
			close(ch)
		}
	}()

	time.Sleep(time.Second)
}

func pipeline() {
	fmt.Println("pipeline")
	ch1 := produce(2)
	ch3 := produce(10)
	ch := merge(ch1, ch3)
	ch = pow2(ch)
	ch = pow2(ch)
	ch = pow2(ch)
	<-consume(ch)
}

func main() {
	//splitFull()
	//pipeline()
	fanOutFull()
}
