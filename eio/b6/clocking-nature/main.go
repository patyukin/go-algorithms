package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		ch <- struct{}{}
		fmt.Println("sending")
	}()

	go func() {
		<-ch
		fmt.Println("received")
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("sent")
}
