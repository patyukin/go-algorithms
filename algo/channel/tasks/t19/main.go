// You can edit this code!
// Click here and start typing.
package main

import (
	"sync"
)

var (
	wg   sync.WaitGroup
	lock sync.Mutex
)

func main() {
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)

	go produce([]int{-6, -3, 0, 3, 7, 9, 10}, ch1)
	go produce([]int{-10, -1, 1, 8, 30, 40, 50}, ch2)

	merge(ch1, ch2)
}

func produce(arr []int, out chan int) {
	for _, v := range arr {
		out <- v
	}
	close(out)
}

func merge(ch1, ch2 chan int) {

}
