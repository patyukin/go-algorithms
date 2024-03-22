package main

import (
	"fmt"
	"time"
)

type Order struct {
	id int
}

func main() {
	orders := make(chan Order, 5)
	cookedOrders := make(chan Order, 5)

	for i := 0; i < 5; i++ {
		orders <- Order{id: i + 1}
	}

	go waiter(orders, cookedOrders)
	go cook(orders, cookedOrders)

	// Sleep to allow other goroutines to finish.
	time.Sleep(10 * time.Second)
}

func waiter(orders, cookedOrders chan Order) {
	for {
		select {
		case order := <-orders:
			fmt.Printf("Waiter received order %d\n", order.id)
		case order := <-cookedOrders:
			fmt.Printf("Waiter served order %d\n", order.id)
		case <-time.After(1 * time.Second):
			fmt.Println("Waiter timed out")
			return
		}
	}
}

func cook(orders, cookedOrders chan Order) {
	for {
		select {
		case order := <-orders:
			fmt.Printf("Cook received order %d, starting to cook...\n", order.id)
			time.Sleep(1 * time.Second) // cooking
			fmt.Printf("Cook finished order %d\n", order.id)
			cookedOrders <- order
		case <-time.After(1 * time.Second):
			fmt.Println("Cook timed out")
			return
		}
	}
}
