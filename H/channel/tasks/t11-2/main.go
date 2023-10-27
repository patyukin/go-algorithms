package main

import (
	"fmt"
	"time"
)

type User struct {
	ID     string
	Amount uint
}

func main() {
	users := []User{
		{"Q", 100},
		{"W", 200},
		{"E", 300},
		{"R", 400},
		{"T", 500},
	}

	requests := make(chan User, len(users))
	result := make(chan string, len(users))

	go atm(requests, result)

	for _, user := range users {
		requests <- user
		<-result
	}
}

func atm(requests chan User, result chan string) {
	for req := range requests {
		fmt.Printf("User %s requested to withdraw $%d\n", req.ID, req.Amount)
		time.Sleep(1 * time.Second)
		fmt.Printf("User %s has successfully withdrawn $%d\n", req.ID, req.Amount)
		result <- req.ID
	}
}
