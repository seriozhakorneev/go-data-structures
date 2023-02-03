// TODO : asdasda
package main

//package workerpool

import (
	"fmt"
	"log"
	"runtime"
)

const (
	//TODO other consts

	// Tasks execution result statuses.
	statusSuccess = iota + 1
	statusFailed
)

type taskType func() taskResult

// TODO: doc
type WorkerPool struct {
	// wCount - count of workers.
	wCount int
	// executeC - a queue channel that delivers tasks to workers.
	executeC chan taskType
	// resultsC - data collection channel from executed workers.
	resultsC chan taskResult

	//TODO: singleton
}

// taskResult - task execution result
type taskResult struct {
	// Status - task status at the end of execution.
	Status int
	// AddInfo - additional information field.
	AddInfo string
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

// Run - runs background workers(goroutines).Count of workers depends
// on workers count field wCount, provided in New. Every worker
// takes task, execute and write result via writeResult.
func (wp *WorkerPool) Run() {
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

// AddTasks - sending tasks to workers through executeC channel,
// if sends more than one, for better performance, should be sent as slice
// (not for loop with repeatedly calling AddTasks).
func (wp *WorkerPool) AddTasks(tasks ...taskType) {
	if tasks == nil {
		return
	}

	go func() {
		for _, task := range tasks {
			wp.executeC <- task
		}
	}()
}

// writeResult - writes result to resultsC channel.
func (wp *WorkerPool) writeResult(result taskResult) {
	wp.resultsC <- result
}

// Result - returns results channel.
func (wp *WorkerPool) Result() chan taskResult {
	return wp.resultsC
}

func main() {
	workers := 12

	wp, err := New(workers)
	if err != nil {
		log.Fatal("failed to create worker pool, %w", err)
	}

	wp.Run()

	wp.AddTasks([]taskType{
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
			return taskResult{Status: statusFailed, AddInfo: "some error description"}
		},
		func() taskResult {
			return taskResult{Status: statusSuccess}
		},
	}...)

	for {
		select {
		case l := <-wp.Result():
			fmt.Println("result:", l.Status, l.AddInfo, "|goroutines:", runtime.NumGoroutine()-2)
		}
	}
}
