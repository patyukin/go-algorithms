package main

import (
	"fmt"
	"sync"
)

type Order struct {
	n               int
	dishList        []string
	deliveryAddress string
	confirm         bool
}

func main() {
	orders := make(chan Order)
	readyOrders := make(chan int)
	drivers := make(chan int)

	var wgDrivers sync.WaitGroup

	wgDrivers.Add(1)

	go func() {
		for o := range orders {
			o.confirm = true
			readyOrders <- o.n
		}
		close(readyOrders)
	}()

	go func() {
		for ro := range readyOrders {
			drivers <- ro
		}
		close(drivers)
	}()

	go func() {
		for d := range drivers {
			fmt.Println("Driver", d, "is processing an order")
		}
		defer wgDrivers.Done()
	}()

	orders <- Order{n: 1, confirm: false, deliveryAddress: "ADDRESS 1"}
	orders <- Order{n: 2, confirm: false, deliveryAddress: "ADDRESS 2"}
	orders <- Order{n: 3, confirm: false, deliveryAddress: "ADDRESS 3"}

	close(orders)
	wgDrivers.Wait()
}
