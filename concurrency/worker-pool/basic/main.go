package main

import (
	"fmt"
	"time"
)

func main() {

	totalWorker := 5
	totalJobs := 5

	jobs := make(chan int)
	result := make(chan int)

	for i := 0; i < totalWorker; i++ {
		go worker(i, jobs, result)
	}

	for i := 0; i < totalJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < totalJobs; i++ {
		fmt.Println(<-result)
	}

}

func worker(workerId int, jobs chan int, result chan int) {

	for job := range jobs {
		fmt.Println("worker id", workerId, "started job", job)
		time.Sleep(time.Second * 2)
		fmt.Println("worker id", workerId, "finished job", job)
		result <- job * 2
	}

}
