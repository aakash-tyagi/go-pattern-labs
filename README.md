# go-pattern-labs

# Go Concurrency & Systems Patterns (Go Labs)

This repository is a **practice-driven learning lab** for mastering **idiomatic Go**, with a strong focus on **concurrency, cancellation, and lifecycle management**.

Instead of solving isolated problems, this repo explores **core Go patterns** that repeatedly appear in real-world backend systems such as high-traffic services, background workers, and cloud-native applications.

The emphasis is on:
- clarity over cleverness
- predictable concurrency
- safe shutdowns
- production-ready thinking

---

## Why This Repository Exists

Most real Go code is not about complex algorithms.  
It is about **controlling work**:

- how much work happens at once  
- when work should stop  
- how failures are handled  
- how systems shut down safely  

This repository exists to practice those fundamentals deliberately and visibly.

---

## Learning Philosophy

Each example in this repository follows a **pattern-first approach**:

> The problem is only a vehicle.  
> The pattern is the lesson.

Every implementation is:
- small and focused
- intentionally constrained
- documented with trade-offs
- written as if it were running in production

---

## Patterns Covered (Phase 1)

This phase focuses on **six foundational Go patterns** that form the backbone of most Go backend systems.

### 1. Worker Pool
📁 `concurrency/worker-pool`

Controls concurrency using a fixed number of workers pulling from a shared job queue.

Used in:
- background workers
- message queue consumers
- batch processing

---

### 2. Fan-out / Fan-in
📁 `concurrency/fanout-fanin`

Executes work in parallel and aggregates results efficiently.

Used in:
- API aggregation
- search queries
- latency-sensitive workloads

---

### 3. Pipeline
📁 `concurrency/pipeline`

Breaks complex workflows into independent processing stages.

Used in:
- log processing
- ETL systems
- streaming pipelines

---

### 4. Backpressure
📁 `concurrency/backpressure`

Protects systems when producers outpace consumers.

Used in:
- event ingestion
- metrics pipelines
- high-traffic systems

---

### 5. Context Propagation
📁 `context/cancellation`

Ensures work can be cancelled safely and predictably.

Used in:
- HTTP servers
- microservices
- CLI tools

---

### 6. Graceful Shutdown
📁 `lifecycle/graceful-shutdown`

Ensures services terminate cleanly without data loss.

Used in:
- Kubernetes workloads
- long-running services
- background processors

---

## Repository Structure

```text
go-patterns-lab/
│
├── concurrency/
│   ├── worker-pool/
│   ├── fanout-fanin/
│   ├── pipeline/
│   └── backpressure/
│
├── context/
│   └── cancellation/
│
├── lifecycle/
│   └── graceful-shutdown/
│
└── README.md
