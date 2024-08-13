package unit_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	input := [...]int{1, 2, 3, 4, 5}
	expected := [...]int{1, 4, 9, 16, 25}
	for i := 0; i < len(input); i++ {
		if Square(input[i]) != expected[i] {
			t.Errorf("Square(%d) = %d, expected %d", input[i], Square(input[i]), expected[i])
		}
	}
}

func TestSquareWithAssert(t *testing.T) {
	input := [...]int{1, 2, 3, 4, 5}
	expected := [...]int{1, 4, 9, 16, 25}
	for i := 0; i < len(input); i++ {
		ret := Square(input[i])
		assert.Equal(t, expected[i], ret)
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("This is an error message")
	fmt.Println("End")
}
func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("This is an error message")
	fmt.Println("End")
}
