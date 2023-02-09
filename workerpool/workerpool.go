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
	// workload - stores the boolean value of workers workload.
	// needed for synchronization.
	workload atomic.Bool
	// executeC - a queue channel that delivers tasks to workers.
	executeC chan DefaultTaskType
	// resultsC - data collection channel from executed workers.
	resultsC chan DefaultTaskResult
}

// DefaultTaskType - default task type.
type DefaultTaskType func() DefaultTaskResult

// DefaultTaskResult - task execution result.
type DefaultTaskResult struct {
	// Status - task status at the end of execution.
	Status int
	// AddInfo - additional information field.
	AddInfo interface{}
}

// New - returns *workerPool with provided count of workers, channels buffer.
// If buffer not provided, channels will be non-buffered.
func New(workers int, buffer ...int) (*WorkerPool, error) {
	if workers < 1 {
		return nil, fmt.Errorf("minimum workers count: 1, got: %d", workers)
	}

	e, r := make(chan DefaultTaskType), make(chan DefaultTaskResult)
	if len(buffer) > 1 {
		e = make(chan DefaultTaskType, buffer[0])
		r = make(chan DefaultTaskResult, buffer[0])
	}

	return &WorkerPool{wCount: workers, executeC: e, resultsC: r}, nil
}

// Run - runs background workers(goroutines).
// Count of workers depends on field wCount, provided in New.
// Every worker takes task, execute and write result via writeResult.
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
func (wp *WorkerPool) AddTask(task DefaultTaskType) (err error) {
	if task == nil {
		return fmt.Errorf("failed to add <nil> task")
	}

	go func() {
		wp.executeC <- task
	}()

	return nil
}

// writeResult - writes result to resultsC channel.
func (wp *WorkerPool) writeResult(result DefaultTaskResult) {
	wp.resultsC <- result
}

// Result - returns results channel.
func (wp *WorkerPool) Result() chan DefaultTaskResult {
	return wp.resultsC
}

// Loaded - returns true if any worker has a task to perform, false if they are all free.
func (wp *WorkerPool) Loaded() bool {
	return wp.workload.Load()
}

// Stop - closing all workers in worker pool, if they are not work loaded.
func (wp *WorkerPool) Stop() error {
	if wp.Loaded() {
		return fmt.Errorf("tasks are in process")
	}

	wp.cancel()

	return nil
}
