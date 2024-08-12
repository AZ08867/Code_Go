package util_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfTasks := 10
	c := make(chan string, numOfTasks)
	for i := 0; i < numOfTasks; i++ {
		go func(i int) {
			ret := runTask(i)
			c <- ret
		}(i)
	}
	finalResult := ""
	for i := 0; i < numOfTasks; i++ {
		finalResult += <-c + "\n"
	}
	return finalResult
}

func TestAllResponse(t *testing.T) {
	t.Log("Before", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Millisecond * 1)
	t.Log("After", runtime.NumGoroutine())
}
