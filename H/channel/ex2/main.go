package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "foo"
	ch <- "bar"
	ch <- "baz" // Послать три значения в буферизованный канал
	close(ch)   // Закрыть канал

	for s := range ch {
		fmt.Println(s)
	}
}
