package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Worker struct {
	id    uint
	name  string
	tasks []string
}

var allTasks = [...]string{"audio", "video", "computer"}

func main() {
	tasks := make(chan string)
	var wg sync.WaitGroup

	worker1 := Worker{id: 1, name: "Petya", tasks: []string{"computer"}}
	worker2 := Worker{id: 1, name: "Vasya", tasks: []string{"video"}}
	worker3 := Worker{id: 1, name: "Ivan", tasks: []string{"audio"}}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			tasks <- allTasks[rand.Intn(3)]
		}
		close(tasks)
	}()

	go func() {
		defer wg.Done()
		for t := range tasks {
			switch t {
			case worker1.tasks[0]:
				fmt.Printf("%s task assigned to %s\n", t, worker1.name)
			case worker2.tasks[0]:
				fmt.Printf("%s task assigned to %s\n", t, worker2.name)
			case worker3.tasks[0]:
				fmt.Printf("%s task assigned to %s\n", t, worker3.name)
			}
		}
	}()

	wg.Wait()
}
