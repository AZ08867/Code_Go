package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("New object created")
			return new(int) // 返回指针类型
		},
	}

	v := pool.Get().(*int) // 获取指针类型
	*v = 100               // 设置值
	fmt.Println(*v)
	pool.Put(v)  // 放回指针类型
	runtime.GC() // 手动触发GC
}

func TestSyncPoolInMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("New object created")
			return new(int) // 返回指针类型
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v := pool.Get().(*int) // 获取指针类型
			*v = 100               // 设置值
			fmt.Println(*v)
			pool.Put(v)  // 放回指针类型
		}()
	}
	wg.Wait()
	runtime.GC() // 手动触发GC
}