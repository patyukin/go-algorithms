package main

import "sync"

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}

		close(a)
	}()

	go func() {
		for _, num := range []int{10, 20, 30} {
			b <- num
		}

		close(b)
	}()

	go func() {
		for _, num := range []int{100, 200, 300} {
			c <- num
		}

		close(c)
	}()

	for num := range joinChannels(a, b, c) {
		println(num)
	}
}

func joinChannels(chs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(chs))

	go func() {
		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for num := range ch {
					out <- num
				}
			}(ch, &wg)
		}

		wg.Wait()
		close(out)
	}()

	return out
}
