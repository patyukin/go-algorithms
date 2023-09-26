package main

import (
	"fmt"
	"time"
)

func generateFibonacci(ch chan int, quite chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quite:
			fmt.Println("stopped")
			return
		}
	}
}

func printFibonacci(ch chan int, quite chan bool) {
	for {
		select {
		case num := <-ch:
			time.Sleep(time.Microsecond * 200000)
			fmt.Println(num)
		case <-time.After(time.Second):
			fmt.Println("Timeout: Took too long to generate next number.")
			quite <- true
			close(ch)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quite := make(chan bool)

	go generateFibonacci(ch, quite)
	go printFibonacci(ch, quite)

	time.Sleep(2 * time.Second)
	quite <- true
	close(ch)
	close(quite)
}
