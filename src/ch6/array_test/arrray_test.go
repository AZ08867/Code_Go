package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var a [5]int
	a[1] = 5
	t.Log(a[1], a[3], a[4])

	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i, v := range arr3 {
		t.Log(i, v)
	}
}

func TestArraySection(t *testing.T) {
	arr4 := [...]int{1, 2, 3, 4, 5}
	arr5 := arr4[1:3]
	t.Log(arr5)
}