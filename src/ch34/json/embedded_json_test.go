package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{"basic_info":{"name":"John","age":30,"address":"New York"},"job_info":{"skills":["Go","Python","Rust"],"company":"ABC","position":"Developer"}}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err != nil {
		t.Errorf("Marshal failed: %v", err)
	} else {
		fmt.Println(string(v))
	}

}
