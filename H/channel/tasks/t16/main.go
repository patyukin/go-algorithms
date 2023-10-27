package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func send(ch chan int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()
	num := rand.Intn(1000)

	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- num
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go send(ch1, 1, &wg)
	go send(ch2, 3, &wg)

	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					fmt.Println("channel ch1 closed")
					break
				}
				fmt.Println("v from channel ch1:", v)
			case v, ok := <-ch2:
				if !ok {
					fmt.Println("channel ch2 closed")
					break
				}
				fmt.Println("v from channel ch2:", v)
			}
		}
	}()

	wg.Wait()
}
