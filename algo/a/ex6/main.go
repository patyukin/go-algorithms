package main

func worker(jobs <-chan int, results chan<- int, multiplier func(int) int) {
	for j := range jobs {
		results <- multiplier(j)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	multiplier := func(x int) int {
		return x * 10
	}

	for w := 1; w <= 3; w++ {
		go worker(jobs, results, multiplier)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		println(<-results)
	}
}
