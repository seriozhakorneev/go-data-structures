// package workerpool
package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
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

// New - returns *workerPool with provided count of workers.
// If workers count of workers not provided it will be set to runtime.NumCPU()-1.
func New(workers ...int) *WorkerPool {

	if len(workers) < 1 {
		// set default workers count
		workers = append(workers, runtime.NumCPU()-1)
	}

	return &WorkerPool{
		wCount:   workers[0],
		executeC: make(chan taskType),
		resultsC: make(chan taskResult),
	}
}

// Run - runs background workers(goroutines).Count of workers depends
// on workers count field wCount, provided in New. Every worker
// takes task, execute and write result via writeResult.
func (wp *WorkerPool) Run() {
	ctx := context.Background()
	ctx, wp.cancel = context.WithCancel(ctx)

	wp.Once.Do(func() {
		log.Printf("Pool started with %d workers\n", wp.wCount)
		wp.startPool(ctx)
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

// AddTasks - sending tasks to workers through executeC channel,
// if sends more than one, for better performance, should be sent as slice,
// not for loop with repeatedly calling AddTasks.
func (wp *WorkerPool) AddTasks(tasks ...taskType) (err error) {
	err = fmt.Errorf("failed to add zero tasks")

	for _, task := range tasks {
		if task == nil {
			return fmt.Errorf("failed to add <nil> task")
		}
		err = nil
	}

	if err != nil {
		return err
	}

	go func() {
		for _, task := range tasks {
			wp.executeC <- task
		}
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
