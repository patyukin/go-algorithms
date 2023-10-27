package main

import (
	"fmt"
	"time"
)

func fetchData(server string, ch chan string) {
	for {
		time.Sleep(1 * time.Second)
		ch <- "Data from " + server
	}
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)

	go fetchData("Server 1", ch1)
	go fetchData("Server 2", ch2)
	go fetchData("Server 3", ch3)

	for {
		select {
		case data := <-ch1:
			fmt.Println(data)
		case data := <-ch2:
			fmt.Println(data)
		case data := <-ch3:
			fmt.Println(data)
		}
	}
}
