package main

import (
	"fmt"
	"time"
)

func send(messages chan string, done chan bool) {
	defer close(messages)
	for i := 1; i <= 2; i++ {
		message := fmt.Sprintf("Message: %d", i)
		messages <- message
		time.Sleep(time.Second)
	}

	done <- true
}

func receive(messages <-chan string, done chan<- bool) {
loop:
	for {
		message, open := <-messages
		if !open {
			break loop
		}

		fmt.Println("Received:", message)
	}

	done <- true
}

func main() {
	messages := make(chan string)
	done := make(chan bool)

	go send(messages, done)
	go receive(messages, done)

	<-done
	<-done
}
