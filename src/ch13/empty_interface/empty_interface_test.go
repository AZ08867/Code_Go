package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer:", i)
	// } else if s, ok := p.(string); ok {
	// 	fmt.Println("String:", s)
	// } else {
	// 	fmt.Println("I don't know what I'm supposed to print!")
	// }

	switch v := p.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("I don't know what I'm supposed to print!")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("hello")
}
