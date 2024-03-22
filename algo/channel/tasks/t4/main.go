package main

import (
	"fmt"
	"sort"
)

func main() {
	var i int
	ch := make(chan int)

	go func() {
		for {
			fmt.Scanln(&i)
			if i == 0 {
				break
			}
			ch <- i
		}
		close(ch)
	}()

	var s []int
	for v := range ch {
		s = append(s, v)
	}

	sort.Ints(s)
	fmt.Println(s)
}
