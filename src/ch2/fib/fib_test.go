package fib

import "testing"

func TestFib(t *testing.T) {
	// var a, b int
	// a, b = 1, 1

	a, b := 0, 1
	for i := 0; i < 10; i++ {
		t.Log(a)
		a, b = b, a+b
	}
}
