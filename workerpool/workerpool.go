package workerpool

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

// Tasks execution result statuses.
const (
	StatusSuccess = iota + 1
	StatusFailed
)

// WorkerPool - a pool of fixed workers(goroutines),
// performing constantly arriving tasks.
type WorkerPool struct {
	sync.Once
	cancel context.CancelFunc

	// wCount - count of workers.
	wCount int
	// workload - stores the boolean value of workload of workers.
	workload atomic.Bool
	// executeC - a queue channel that delivers tasks to workers.
	executeC chan taskType
	// resultsC - data collection channel from executed workers.
	resultsC chan taskResult
}

// taskType - default task type.
type taskType func() taskResult

// taskResult - task execution result.
type taskResult struct {
	// Status - task status at the end of execution.
	Status int
	// AddInfo - additional information field.
	AddInfo interface{}
}

// New - returns *workerPool with provided count of workers, channels buffer.
// If buffer not provided, channels will be non-buffered.
func New(workers int, buffer ...int) *WorkerPool {
	e, r := make(chan taskType), make(chan taskResult)

	if len(buffer) > 1 {
		e = make(chan taskType, buffer[0])
		r = make(chan taskResult, buffer[0])
	}

	return &WorkerPool{wCount: workers, executeC: e, resultsC: r}
}

// Run - runs background workers(goroutines).Count of workers depends
// on workers count field wCount, provided in New. Every worker
// takes task, execute and write result via writeResult.
func (wp *WorkerPool) Run() {
	ctx := context.Background()
	ctx, wp.cancel = context.WithCancel(ctx)

	wp.Once.Do(func() {
		wp.startPool(ctx)
		log.Printf("Pool started with %d workers\n", wp.wCount)
	})
}

// startPool - spawns workers-goroutines, make them listening to incoming tasks.
func (wp *WorkerPool) startPool(ctx context.Context) {
	for i := 0; i < wp.wCount; i++ {
		go func() {
			for {
				select {
				case task := <-wp.executeC:
					_ = wp.workload.Swap(true)
					// Executing task and writing its results.
					wp.writeResult(task())
					_ = wp.workload.Swap(false)
				case <-ctx.Done():
					// finish worker
					return
				default:
					continue
				}
			}
		}()
	}
}

// AddTask - sending task to workers through executeC channel.
func (wp *WorkerPool) AddTask(tasks taskType) (err error) {
	if tasks == nil {
		return fmt.Errorf("failed to add <nil> task")
	}

	go func() {
		wp.executeC <- tasks
	}()

	return nil
}

// writeResult - writes result to resultsC channel.
func (wp *WorkerPool) writeResult(result taskResult) {
	wp.resultsC <- result
}

// Result - returns results channel.
func (wp *WorkerPool) Result() chan taskResult {
	return wp.resultsC
}

// Loaded - returns true if any worker has a task to perform, false if they are all free.
func (wp *WorkerPool) Loaded() bool {
	return wp.workload.Load()
}

// Stop - closing all workers in worker pool, if they are not work loaded.
func (wp *WorkerPool) Stop() error {
	if wp.Loaded() {
		return fmt.Errorf("some tasks are in process")
	}

	wp.cancel()

	return nil
}
