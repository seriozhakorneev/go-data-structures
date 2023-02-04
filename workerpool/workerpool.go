// package workerpool
package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

const (
	//TODO other consts

	// Tasks execution result statuses.
	statusSuccess = iota + 1
	statusFailed
)

// taskType - default task type.
type taskType func() taskResult

// WorkerPool - a pool of fixed workers(goroutines),
// performing constantly arriving tasks.
type WorkerPool struct {
	sync.Once
	// wCount - count of workers.
	wCount int
	// executeC - a queue channel that delivers tasks to workers.
	executeC chan taskType
	// resultsC - data collection channel from executed workers.
	resultsC chan taskResult
}

// taskResult - task execution result.
type taskResult struct {
	// Status - task status at the end of execution.
	Status int
	// AddInfo - additional information field.
	AddInfo interface{}
}

// New - returns *workerPool with provided count of workers.
func New(workers int) (*WorkerPool, error) {
	if workers < 1 {
		return nil, fmt.Errorf(
			"workers count should be more than 1, got: %d",
			workers,
		)
	}

	return &WorkerPool{
		wCount:   workers,
		executeC: make(chan taskType),
		resultsC: make(chan taskResult),
	}, nil
}

// startPool - spawns workers-goroutines, make them listening to incoming tasks
func (wp *WorkerPool) startPool() {
	for ; wp.wCount > 0; wp.wCount-- {
		go func() int {
			for {
				select {
				case task := <-wp.executeC:
					wp.writeResult(task())
				default:
					continue
				}
			}
		}()
	}
}

// Run - runs background workers(goroutines).Count of workers depends
// on workers count field wCount, provided in New. Every worker
// takes task, execute and write result via writeResult.
func (wp *WorkerPool) Run() {
	wp.Once.Do(func() {
		log.Printf("Worker pool started with %d workers\n", wp.wCount)
		wp.startPool()
	})
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

// TODO: method stop worker pool with quit channel
// TODO: method to let user know if there tasks still spinning in wp

func main() {
	workers := 12

	wp, err := New(workers)
	if err != nil {
		log.Fatal("failed to create worker pool: ", err)
	}

	wp.Run()

	err = wp.AddTasks([]taskType{
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess, AddInfo: 12 * 356}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusFailed, AddInfo: "some error description"}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess, AddInfo: 12 * 356}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusFailed, AddInfo: "some error description"}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
	}...)
	if err != nil {
		log.Fatal("failed to add tasks: ", err)
	}

	for {
		select {
		case l := <-wp.Result():
			fmt.Println("result:", l.Status, l.AddInfo, "|goroutines:", runtime.NumGoroutine())
		}
	}
}
