package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 2, 4}
	t.Log(a == b, a == c, b == c)
}

func TestBitClear(t *testing.T) {
	a := 7 // 0111 in binary
	a = a &^ Writable

	b := 1 // 0001 in binary
	t.Log(a&Readable, a&Writable, a&Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(b&Readable, b&Writable, b&Executable)
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
}
