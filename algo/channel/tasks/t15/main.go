package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- fmt.Sprintf("#%d", i)
		}
		close(ch)
	}()

	go func() {
		timeout := time.After(5 * time.Second)
		for {
			select {
			case data := <-ch:
				fmt.Println("data = ", data)
			case data := <-ch:
				fmt.Println("data = ", data)
			case data := <-ch:
				fmt.Println("data = ", data)
			case data := <-ch:
				fmt.Println("data = ", data)
			case data := <-ch:
				fmt.Println("data = ", data)
			default:
				fmt.Println("timeout = ", timeout)
			}
		}
	}()

	wg.Wait()
}
