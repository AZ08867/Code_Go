package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultipleValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultipleValues()
	t.Log(a, b)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func TestReturnMultipleValues(t *testing.T) {
	a, b := returnMultipleValues()
	t.Log(a, b)
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

// defer 函数
func Clear() {
	fmt.Println("clear resources")
}
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	// panic("err")
}
