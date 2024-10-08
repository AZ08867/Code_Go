package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	t.Log("CPTestService")
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	ch := make(chan string)
	go func() {
		fmt.Println("goroutine started")
		ch <- service()
		fmt.Println("service exited")
	}()
	return ch
}

func TestAsyncService(t *testing.T) {
	t.Log("CPTestAsyncService")
	ch := AsyncService()
	otherTask()
	fmt.Println(<-ch)
	fmt.Println("over")
}
