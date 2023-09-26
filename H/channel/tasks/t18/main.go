package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID       int
	clientID int
}

type Cook struct {
	ID      int
	orderID int
}

type ReadyOrder struct {
	ID      int
	orderID int
}

func main() {
	order := make(chan Order)
	cookChannel := make(chan Cook)
	readyOrder := make(chan ReadyOrder)

	go func() {
		defer close(order)
		for i := 1; i <= 5; i++ {
			order <- Order{ID: i, clientID: rand.Intn(100)}
		}
	}()

	go func() {
		defer close(cookChannel)
		i := 0
		for o := range order {
			i++
			cookChannel <- Cook{ID: i, orderID: o.ID}
		}
	}()

	go func() {
		defer close(readyOrder)
		i := 0
		for c := range cookChannel {
			i++
			time.Sleep(1 * time.Second)
			readyOrder <- ReadyOrder{ID: i, orderID: c.orderID}
		}
	}()

	for ro := range readyOrder {
		fmt.Printf("read order #%d\n", ro.orderID)
	}
}
