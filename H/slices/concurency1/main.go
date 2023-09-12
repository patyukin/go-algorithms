package main

import (
	"fmt"
	"sync"
)

func main() {
	unsafeWrites()
}

func unsafeWrites() {
	slice := []int{0, 0, 0, 0}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		for index := range slice {
			index := index
			wg.Add(1)

			go func() {
				defer wg.Done()
				slice[index]++
			}()
		}
	}

	wg.Wait()

	fmt.Println(slice)
}
