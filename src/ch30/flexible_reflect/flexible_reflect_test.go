package flexible_reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type Customer struct {
	Name       string
	Age        int
	CustomerID string
}

type Employee struct {
	Name       string `format:"normal"`
	Age        int
	EmployeeID string
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "teo", 3: "three"}
	b := map[int]string{1: "one", 2: "teo", 4: "three"}
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{Name: "Jack", Age: 30, CustomerID: "12345"}
	c2 := Customer{"Jack", 30, "12345"}
	t.Log("c1 == c2?", reflect.DeepEqual(c1, c2))
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		// Elem() 获取指针指向的元素
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return fmt.Errorf("Only accept struct or pointer to struct")
		}
	}
	if settings == nil {
		return errors.New("No settings provided")
	}

	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		if field, ok = (reflect.TypeOf(st)).Elem().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			v_str := reflect.ValueOf(st)
			v_str = v_str.Elem()
			v_str.FieldByName(k).Set(reflect.ValueOf(v))

		}
	}
	return nil

}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Jack", "Age": 45}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)

}
