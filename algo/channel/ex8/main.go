package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {
	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

//func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
//	out := make(chan int)
//	go func() {
//		defer close(out)
//		select {
//		case val := <-firstChan:
//			out <- val * val
//		case val := <-secondChan:
//			out <- val * 3
//		case <-stopChan:
//			return
//		}
//	}()
//
//	return out
//}

func calc(fn func(int) int, in <-chan int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- fn(<-in)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			x1 := <-in1
			x2 := <-in2

			f1 := make(chan int, 1)
			f2 := make(chan int, 1)

			go func() { f1 <- fn(x1) }()
			go func() { f2 <- fn(x2) }()

			out <- <-f1 + <-f2
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)
	var sum int
	go func() {
		defer close(out)
		select {
		case v := <-arguments:
			sum += v
		case <-done:
			out <- sum
			return
		}
	}()

	return out
}
