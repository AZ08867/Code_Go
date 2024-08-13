package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }
	for i := 0; i < 11; i++ {
		obj, err := pool.GetObj(time.Second * 1)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Printf("Get obj %d\n", i)
			if err := pool.ReleaseObj(obj); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("PASS or Done")
}
