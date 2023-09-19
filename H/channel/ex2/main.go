package main

import (
	"fmt"
	"math/rand"
	"time"
)

func client(clientToServer chan<- int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 5; i++ {
		number := rand.Intn(10)
		fmt.Println(number)
		clientToServer <- number
	}

	close(clientToServer)
}

func server(clientToServer chan int, serverToClient chan<- int) {
	sum := 0

	for v := range clientToServer {
		sum += v
	}

	serverToClient <- sum
	close(serverToClient)
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go client(in)
	go server(in, out)

	sum := <-out

	fmt.Println("Сумма чисел:", sum)
}
