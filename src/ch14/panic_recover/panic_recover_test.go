package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recovered:", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("error, Something went wrong"))
}
