package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"huihui.com",
	"google.com",
	"example.com",
	"baddomain.com",
	"huihui.com",
	"huihui.com",
	"google.com",
	"example.com",
	"baddomain.com",
	"huihui.com",
}

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	totalWorker := 5
	wg := sync.WaitGroup{}

	jobs := make(chan string)
	result := make(chan HealthResult)

	for i := 1; i <= totalWorker; i++ {
		go worker(ctx, i, &wg, jobs, result)
	}

	go func() {
		for i := 0; i < len(urls); i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- urls[i]:
				wg.Add(1)
			}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	badUrl := 0
	for res := range result {
		if !res.Ok {
			badUrl++
		}
		if badUrl > 2 {
			fmt.Println("3 bad urls, cancelling execution")
			cancel()
		}

		fmt.Println(res)
	}
}

type HealthResult struct {
	Url string
	Ok  bool
}

func worker(ctx context.Context, workerId int, wg *sync.WaitGroup, jobs chan string, result chan HealthResult) {

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	for {
		select {
		case <-ctx.Done():
			return

		case url, ok := <-jobs:
			if !ok {
				return
			}

			fmt.Println("worker id ", workerId, " picked url ", url)
			okUrl := CheckURL(ctx, client, url)
			fmt.Println("worker id ", workerId, " checked url ", url)
			result <- HealthResult{
				Url: url,
				Ok:  okUrl,
			}
			wg.Done()
		}
	}
}

func CheckURL(ctx context.Context, client http.Client, url string) bool {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://"+url,
		nil,
	)

	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK

}
