package main

import (
	"errors"
	"fmt"
	"log"
)

type Worker struct {
	workerID uint
	done     bool
}

type Task struct {
	taskID uint
}

type Result struct {
	workerID uint
	taskID   uint
}

func main() {
	var workers []Worker
	var tasks []Task

	for i := 1; i < 50; i++ {
		workers = append(workers, Worker{uint(i), false})
	}

	for i := 5; i < 10; i++ {
		tasks = append(tasks, Task{uint(i)})
	}

	taskResults := make(chan Result)

	go func() {
		for _, t := range tasks {
			w, err := getReadWorker(workers)
			if err != nil {
				log.Fatalf("erorr: %v", err)
			}

			taskResults <- Result{
				workerID: w.workerID,
				taskID:   t.taskID,
			}
		}

		close(taskResults)
	}()

	for v := range taskResults {
		fmt.Println(v)
	}
}

func getReadWorker(workers []Worker) (Worker, error) {
	for _, w := range workers {
		if !w.done {
			return w, nil
		}
	}

	return Worker{}, errors.New("sdfsdfsdf")
}
