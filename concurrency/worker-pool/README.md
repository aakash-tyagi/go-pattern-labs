# Worker Pool Pattern (Go)

The **Worker Pool** pattern limits concurrency by running a fixed number of workers that pull jobs from a shared queue.

This is one of the **most important patterns in Go**, used whenever:
- work must be bounded
- resources are limited
- spawning unbounded goroutines is unsafe

---

## Problem This Pattern Solves

Without a worker pool:
- one goroutine per task → memory spikes
- uncontrolled parallelism → CPU & network exhaustion
- hard-to-cancel work

The worker pool introduces **backpressure and predictability**.

---

## Core Idea

- A **jobs channel** holds incoming work
- A fixed number of **workers** pull from the channel
- Results are sent to a **results channel**
- Workers exit cleanly when the job channel is closed

---

## Minimal Example

```go
jobs := make(chan Job)
results := make(chan Result)

for i := 0; i < workerCount; i++ {
    go worker(jobs, results)
}

for _, job := range input {
    jobs <- job
}
close(jobs)

for range input {
    <-results
}
