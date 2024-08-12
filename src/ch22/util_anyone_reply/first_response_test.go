package util_anyone_reply

import (
	"fmt"
	"testing"
	"time"
	"runtime"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfTasks := 10
	c := make(chan string, numOfTasks)
	for i := 0; i < numOfTasks; i++ {
		go func(i int) {
			ret := runTask(i)
			c <- ret
		}(i)
	}
	return <-c
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Millisecond * 1)
	t.Log("After", runtime.NumGoroutine())
}
