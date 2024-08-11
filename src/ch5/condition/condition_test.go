package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1; a == 1 {
		t.Log("a == 1")
	} else if a == 2 {
		t.Log("a == 2")
	} else {
		t.Log("a == 3")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		}
	}
}
