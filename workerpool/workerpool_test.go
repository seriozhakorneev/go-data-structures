package workerpool

import (
	"fmt"
	"testing"
	"time"
)

func TestBench(t *testing.T) {

	tWP := New(100)
	tWP.Run()

	taskCount := 100
	taskDuration := 1 * time.Second

	var err error

	start := time.Now()

	for i := 0; i < taskCount; i++ {
		j := i + 1

		err = tWP.AddTask(func() taskResult {
			// imitate some work
			time.Sleep(taskDuration)
			return taskResult{
				Status:  StatusSuccess,
				AddInfo: j,
			}
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for {
		select {
		case _ = <-tWP.Result():

			taskCount--

			if taskCount == 0 {
				for tWP.Loaded() {
					continue
				}

				err = tWP.Stop()
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(time.Since(start))
				return
			}
		}
	}
}
