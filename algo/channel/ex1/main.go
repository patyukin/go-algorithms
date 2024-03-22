package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string) // Создать канал для передачи строк
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "ping"      // Послать "ping"
		fmt.Println(<-ch) // "pong"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		message := <-ch      // Заблокируется при приеме; запишет строку в message
		fmt.Println(message) // "ping"
		ch <- "pong"         // Заблокируется при передаче
	}()

	wg.Wait()
}
