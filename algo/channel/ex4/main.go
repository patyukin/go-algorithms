package main

import (
	"fmt"
)

func main() {
	firstCh, secondCh := make(chan float64), make(chan float64)

	go func() {
		var first float64
		fmt.Scanln(&first)
		firstCh <- first
		close(firstCh)
	}()

	go func() {
		var second float64
		fmt.Scanln(&second)
		secondCh <- second
		close(secondCh)
	}()

	f := <-firstCh
	s := <-secondCh

	res := (f + s) / 2
	fmt.Println(res)
}
