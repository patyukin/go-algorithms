package main

func split(ch <-chan int) (<-chan int, <-chan int) {
	even := make(chan int)
	odd := make(chan int)

	go func() {
		for v := range ch {
			if v%2 == 0 {
				even <- v
			} else {
				odd <- v
			}
		}

		close(even)
		close(odd)
	}()

	return even, odd
}

func fanOut(in <-chan int, n int) []chan int {
	var chs []chan int
	for i := 0; i < n; i++ {
		ch := make(chan int)
		chs = append(chs, make(chan int))

		go func() {
			ch := ch
			for v := range in {
				ch <- v
			}
			close(ch)
		}()
	}

	return chs
}
