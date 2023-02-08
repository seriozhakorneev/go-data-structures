package workerpool

import (
	"fmt"
	"testing"
	"time"
)

const (
	taskCount        = 10
	taskExecDuration = 1 * time.Second

	minWCount = 2
	maxWCount = 10
)

func BenchmarkWP(b *testing.B) {
	for w := minWCount; w <= maxWCount; w++ {

		b.Run(fmt.Sprintf("input %d", w), func(b *testing.B) {

			wp, err := New(w)
			if err != nil {
				b.Fatalf("failed to make new pool: %s", err)
				return
			}

			wp.Run()

			for t := 0; t < taskCount; t++ {
				j := t + 1
				err = wp.AddTask(func() DefaultTaskResult {
					time.Sleep(taskExecDuration) // imitate some work

					return DefaultTaskResult{
						Status:  StatusSuccess,
						AddInfo: j,
					}
				})
				if err != nil {
					b.Fatalf("failed to add task: %s", err)
					return
				}
			}

			counter := 0

		infinity:
			for {
				select {
				case <-wp.Result():
					counter++

					if counter == taskCount {
						for wp.Loaded() {
							continue
						}

						break infinity
					}
				}
			}

			err = wp.Stop()
			if err != nil {
				b.Fatalf("failed to stop pool: %s", err)
				return
			}
		})
	}
}

func BenchmarkWPBuffered(b *testing.B) {
	for w := minWCount; w <= maxWCount; w++ {

		b.Run(fmt.Sprintf("input %d", w), func(b *testing.B) {

			wp, err := New(w, taskCount)
			if err != nil {
				b.Fatalf("failed to make new pool: %s", err)
				return
			}

			wp.Run()

			for t := 0; t < taskCount; t++ {
				j := t + 1
				err = wp.AddTask(func() DefaultTaskResult {
					time.Sleep(taskExecDuration) // imitate some work

					return DefaultTaskResult{
						Status:  StatusSuccess,
						AddInfo: j,
					}
				})
				if err != nil {
					b.Fatalf("failed to add task: %s", err)
					return
				}
			}

			counter := 0

		infinity:
			for {
				select {
				case <-wp.Result():
					counter++

					if counter == taskCount {
						for wp.Loaded() {
							continue
						}

						break infinity
					}
				}
			}

			err = wp.Stop()
			if err != nil {
				b.Fatalf("failed to stop pool: %s", err)
				return
			}
		})
	}
}
