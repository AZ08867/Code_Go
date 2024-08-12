// Bug 修复：将 select 修改成正确的包名
package my_select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
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

func TestSelect(t *testing.T) {
	t.Log("CPTestSelect")
	select {
	case ret := <-AsyncService():
		fmt.Println(ret)
	case <-time.After(time.Millisecond * 100):
		fmt.Println("time out")
	}
}
