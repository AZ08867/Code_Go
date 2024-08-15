package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.TypeOf(f))
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12.3
	CheckType(f)
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"` // 添加结构体标签
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{Name: "Jack", Age: 30, EmployeeID: "12345"}
	// 按名字获取成员
	nameField := reflect.ValueOf(e).Elem().FieldByName("Name")

	t.Logf("Name: value(%[1]v), type(%[1]T)\n", nameField) // 获取成员类型
	if nameField, ok := reflect.TypeOf(e).Elem().FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format:", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(40)})
	t.Log("Updated Age:", e)
}
