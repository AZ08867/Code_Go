package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct{}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Creating single instance now.")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			// fmt.Println(obj)
			// Bug 修复：将 unsafe.Pointer(obj) 转换为字符串类型，然后输出
			fmt.Printf("%p\n", unsafe.Pointer(obj))
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
