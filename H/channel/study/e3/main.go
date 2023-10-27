package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Job struct {
	id     int
	random int
}

type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	num := number
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.random)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func main() {
	// Creating workers
	createWorkerPool(10)

	// Populate jobs
	for i := 0; i < 10; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)

	// Printing results
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.random, result.sumOfDigits)
	}
}
