package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"1", "Bob", 20}
	e1 := Employee{Name: "Bob", Age: 20}
	e2 := new(Employee)
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", &e)
	t.Logf("e: %v, e1: %v, e2: %v", e, e1, e2)
}

func (e Employee) String() string {
	fmt.Println("Address of e is", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 更推荐使用指针的方式
// func (e *Employee) String() string {
// 	fmt.Println("Address of e is", unsafe.Pointer(&e.Name))

// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Println("Address of e is", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
