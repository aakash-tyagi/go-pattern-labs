# Concurrent URL Health Checker (Worker Pool)

## Problem Statement

We need to check the health of a list of URLs concurrently.

A naive approach of spawning one goroutine per URL does not scale and can
exhaust system resources (CPU, memory, file descriptors).

This problem demonstrates how to solve the task using the **Worker Pool**
concurrency pattern in Go.

---

## Requirements

- Limit the maximum number of concurrent HTTP requests
- Do NOT spawn one goroutine per URL
- Each request must timeout after 2 seconds
- Cancel all in-flight requests if more than 3 URLs fail
- Ensure no goroutine leaks
- All workers must exit cleanly
- Return results as a map: `map[string]bool`

---

## Why Worker Pool?

The worker pool pattern provides:
- bounded concurrency
- predictable resource usage
- controlled shutdown
- backpressure via channels

This mirrors real production systems such as background workers and queue consumers.

---

## High-Level Design

1. A fixed number of workers pull URLs from a `jobs` channel
2. Each worker performs an HTTP request with context
3. Results are sent back to main flow.
