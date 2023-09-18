package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// Обработка задачи
		fmt.Println("Worker", id, "started job", j)
		result := j * 2
		fmt.Println("Worker", id, "finished job", j)

		// Отправка результата
		results <- result
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	// Инициализация каналов для задач и результатов
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Запуск worker'ов
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(w)
	}

	// Отправка задач в канал задач
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Получение результатов из канала результатов
	wg.Wait()
	close(results)
	for r := range results {
		fmt.Println("Result:", r)
	}
}
