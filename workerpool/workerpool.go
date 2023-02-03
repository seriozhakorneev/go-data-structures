// package main
package workerpool

//TODO: func(int) int type
// other consts

// TODO: doc
type WorkerPool struct {
	// wCount - count of workers.
	wCount int
	// executeC - a queue channel that delivers tasks to workers.
	executeC chan func(int) int
	// resultsC - data collection channel from executed workers.
	resultsC chan int

	//TODO: singleton
}

// TODO: doc
func New(workers int) *WorkerPool {
	return &WorkerPool{
		wCount:   workers,
		executeC: make(chan func(int) int),
		resultsC: make(chan int),
	}
}

// TODO: doc
func (wp *WorkerPool) Run() {
	// TODO: убрать i за ненужностью
	for i := 0; i < wp.wCount; i++ {
		go func(workerID int) int {
			for {
				select {
				case task := <-wp.executeC:
					//fmt.Println("worker:", workerID, "taks:", task(workerID))
					wp.writeResult(task(workerID))
				default:
					continue
				}
			}
		}(i + 1)
	}
}

// AddTasks - sending tasks to workers through executeC channel,
// if sends more than one, for better performance, should be sent as slice.
func (wp *WorkerPool) AddTasks(tasks ...func(int) int) {
	//TODO: waitgroup and return ok?
	go func() {
		for _, task := range tasks {
			wp.executeC <- task
		}
	}()
}

// writeResult - writes result from worker execution to resultsC.
func (wp *WorkerPool) writeResult(result int) {
	wp.resultsC <- result
}

// TODO: doc
func (wp *WorkerPool) Result() chan int {
	return wp.resultsC
}

//func main() {
//	workers := 2
//
//	wp := New(workers)
//	wp.Run()
//
//	wp.AddTasks([]func(i int) int{
//		func(i int) int {
//			return 298
//		},
//		func(i int) int {
//			return 400
//		},
//		func(i int) int {
//			return 1892
//		},
//		func(i int) int {
//			return 4892
//		},
//		func(i int) int {
//			return 2
//		},
//	}...)
//
//	wp.AddTasks(func(i int) int {
//		return 5000
//	})
//
//	wp.AddTasks(func(i int) int {
//		return 15000
//	})
//
//	//TODO утащить в функцию для считывания?
//	for {
//		select {
//		case l := <-wp.Result():
//			fmt.Println("result:", l, "goroutines:", runtime.NumGoroutine()-2)
//		}
//	}
//}
