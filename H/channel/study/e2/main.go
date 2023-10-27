package main

func main() {
	ch := make(chan int, 4)

	ch <- 10

	_ = <-ch

	go func() {
		for {
			ch <- 10
		}
	}()

	<-ch

	return
}
