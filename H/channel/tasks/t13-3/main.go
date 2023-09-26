package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData(serverName string, ch chan string) {
	i := 0
	for {
		i++
		ch <- fmt.Sprintf("server: %s send log #%d\n", serverName, i)
		time.Sleep(5 * time.Second)
	}
}

func printLog(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		}
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(3)

	go fetchData("Server 1", ch1)
	go fetchData("Server 2", ch2)
	go fetchData("Server 3", ch3)

	go printLog(ch1, &wg)
	go printLog(ch2, &wg)
	go printLog(ch3, &wg)

	wg.Wait()
}
