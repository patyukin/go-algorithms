package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	ID int
}

func messageSource(messages chan<- Message, n int) {
	for i := 0; i < n; i++ {
		messages <- Message{ID: i}
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
	close(messages)
}

func messageHandler(messages <-chan Message, result chan<- int, expired chan<- int, timeout time.Duration) {
	for message := range messages {
		select {
		case <-time.After(timeout):
			expired <- message.ID
		default:
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			result <- message.ID
		}
	}
	close(result)
	close(expired)
}

func runSystem(numMessages int, processTimeout time.Duration) (int, int) {
	messages := make(chan Message, 5)
	result := make(chan int)
	expired := make(chan int)

	go messageSource(messages, numMessages)
	go messageHandler(messages, result, expired, processTimeout)

	var processed, expiredCount int
	for processed+expiredCount < numMessages {
		select {
		case <-result:
			processed++
		case <-expired:
			expiredCount++
		}
	}

	return processed, expiredCount
}

func main() {
	rand.Seed(time.Now().Unix())
	processed, expired := runSystem(20, 500*time.Millisecond)
	fmt.Printf("Processed messages: %d\nExpired messages: %d\n", processed, expired)
}
