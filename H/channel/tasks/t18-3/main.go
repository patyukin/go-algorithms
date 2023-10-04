package main

import (
	"context"
	"fmt"
	"sync"
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

func waiter(ctx context.Context, order <-chan Order, cookChannel chan<- Cook, readyOrder chan ReadyOrder, wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()
	i := 1
	var owg sync.WaitGroup

	owg.Add(1)
	go func() {
		defer owg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("waiter was canceled.")
				return
			case o, ok := <-readyOrder:
				if !ok {
					return
				}

				fmt.Printf("ready order %v\n", o)
			case <-time.After(3 * time.Hour):
				fmt.Println("Превышен лимит времени! Процесс заказа прерван.")
				cancel()
				return
			}
		}
	}()

	owg.Add(1)
	go func() {
		defer owg.Done()
		defer close(cookChannel)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("waiter was canceled.")
				return
			case o, ok := <-order:
				if !ok {
					return
				}

				time.Sleep(1 * time.Second)
				cookChannel <- Cook{ID: o.ID, orderID: i}
				i++
			case <-time.After(3 * time.Hour):
				fmt.Println("Превышен лимит времени! Процесс заказа прерван.")
				cancel()
				return
			}
		}
	}()

	owg.Wait()
}

func chef(ctx context.Context, cookChannel <-chan Cook, readyOrder chan ReadyOrder, wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()
	defer close(readyOrder)
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("chef was canceled.")
			return
		case co, ok := <-cookChannel:
			if !ok {
				return
			}

			time.Sleep(1 * time.Second)
			ro := ReadyOrder{ID: co.ID, orderID: i}
			fmt.Println("ready order form chef", ro)
			readyOrder <- ro
			i++
		case <-time.After(3 * time.Hour):
			fmt.Println("Превышен лимит времени! Процесс заказа прерван.")
			cancel()
			return
		}
	}
}

func client(ctx context.Context, order chan Order, wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("client was canceled.")
			return
		case order <- Order{ID: i, clientID: i + 100}:
		case <-time.After(3 * time.Hour):
			fmt.Println("Превышен лимит времени! Процесс заказа прерван.")
			cancel()
			return
		}
	}
	close(order)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	order := make(chan Order)
	cookChannel := make(chan Cook)
	readyOrder := make(chan ReadyOrder)

	var wg sync.WaitGroup

	wg.Add(1)
	go client(ctx, order, &wg, cancel)

	wg.Add(1)
	go waiter(ctx, order, cookChannel, readyOrder, &wg, cancel)

	wg.Add(1)
	go chef(ctx, cookChannel, readyOrder, &wg, cancel)

	wg.Wait()
}
