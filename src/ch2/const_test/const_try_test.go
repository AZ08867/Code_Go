package const_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}

func TestConst1(t *testing.T) {
	a := 7 // 0111 in binary
	b := 1 // 0001 in binary
	t.Log(a&Readable, a&Writable, a&Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(b&Readable, b&Writable, b&Executable)
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
}
