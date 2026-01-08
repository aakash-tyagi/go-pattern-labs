package main

import (
	"fmt"
	"time"
)

var urls = []string{
	"google.com",
	"example.com",
	"baddomain.com",
}

func main() {

	totalWorker := 5

	jobs := make(chan string)
	result := make(chan map[string]bool)

	for i := 1; i <= totalWorker; i++ {
		go worker(i, jobs, result)
	}

	for i := 0; i < len(urls); i++ {
		jobs <- urls[i]
	}
	close(jobs)

	for i := 1; i <= len(urls); i++ {
		fmt.Println(<-result)
	}

}

func worker(workerId int, jobs chan string, result chan map[string]bool) {

	for url := range jobs {
		fmt.Println("worker id ", workerId, " picked url ", url)

		// call http
		time.Sleep(2 * time.Second)

		fmt.Println("worker id ", workerId, " checked url ", url)
		result <- map[string]bool{
			url: true,
		}
	}
}
