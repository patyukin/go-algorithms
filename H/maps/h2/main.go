package main

import (
	"fmt"
)

type Message struct {
	operationType string
	data          string
}

func doSomething(ch chan Message) {
	select {
	case msg := <-ch:
		if msg.operationType == "Print" {
			fmt.Println("Received message for Print:", msg.data)
		} else if msg.operationType == "Log" {
			fmt.Println("Received message for Log:", msg.data)
		} else {
			fmt.Println("Received unknown message:", msg.data)
		}
	}
}

func main() {
	ch := make(chan Message)

	go doSomething(ch)

	msg1 := Message{operationType: "Print", data: "Hello, World!"}
	msg2 := Message{operationType: "Log", data: "This is a log message"}

	ch <- msg1
	ch <- msg2
}
