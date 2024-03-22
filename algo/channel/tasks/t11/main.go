package main

import (
	"fmt"
)

type User struct {
	id uint
}

func main() {
	withdraw := make(chan uint)
	show := make(chan int)

	users := []User{{1}, {2}, {3}, {4}}

	go func() {
		for _, user := range users {
			withdraw <- user.id
		}
		close(withdraw)
	}()

	go func() {
		for v := range withdraw {
			show <- int(v)
		}
		close(show)
	}()

	for v := range show {
		fmt.Println(v)
	}
}
