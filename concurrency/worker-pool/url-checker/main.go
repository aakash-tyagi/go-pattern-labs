package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"google.com",
	"example.com",
	"baddomain.com",
	"huihui.com",
}

func main() {

	totalWorker := 5
	wg := sync.WaitGroup{}

	jobs := make(chan string)
	result := make(chan HealthResult)

	for i := 1; i <= totalWorker; i++ {
		go worker(i, &wg, jobs, result)
	}

	wg.Add(len(urls))

	go func() {
		for i := 0; i < len(urls); i++ {

			jobs <- urls[i]
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println(res)
	}
}

type HealthResult struct {
	Url string
	Ok  bool
}

func worker(workerId int, wg *sync.WaitGroup, jobs chan string, result chan HealthResult) {

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	for url := range jobs {

		fmt.Println("worker id ", workerId, " picked url ", url)
		ok := CheckURL(client, url)
		fmt.Println("worker id ", workerId, " checked url ", url)
		result <- HealthResult{
			Url: url,
			Ok:  ok,
		}
		wg.Done()
	}
}

func CheckURL(client http.Client, url string) bool {
	resp, err := client.Get("https://" + url)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK

}
