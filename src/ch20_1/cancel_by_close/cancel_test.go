package cancel_by_close_test

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel_2(ch chan struct{}) {
	close(ch)
}

func TestCancel(t *testing.T) {
	ch := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancelled(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ch)
	}
	// cancel_1(ch)
	cancel_2(ch)
	time.Sleep(time.Second * 1)
}
