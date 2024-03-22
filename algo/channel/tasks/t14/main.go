package main

import (
	"fmt"
)

func process(sendTask, result chan Task) {
	for s := range sendTask {
		s.done = true
		result <- s
	}
	close(result)
}

type Task struct {
	ID   uint
	Name string
	done bool
}

func main() {
	sendTask := make(chan Task)
	result := make(chan Task)

	t1 := Task{ID: 1, Name: "A", done: false}

	go process(sendTask, result)

	sendTask <- t1
	close(sendTask)

	for v := range result {
		fmt.Println("v =", v)
	}
}
