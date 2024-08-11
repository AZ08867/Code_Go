package error_test

import (
	"fmt"
	"strconv"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n must be a positive integer")
	}
	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib, nil
}

func TestGetFibonacci(t *testing.T) {
	fib, err := GetFibonacci(-1)
	if err != nil {
		t.Log(err)
	}
	t.Log(fib)
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	list, err = GetFibonacci(i)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}
