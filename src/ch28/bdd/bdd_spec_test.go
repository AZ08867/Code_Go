package bdd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey" // 引入Convey，因为前面添加了. 所以这就意味着convey中的所有函数是包括于当前命名空间的，可以直接使用
)

func TestSpec(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given two even numbers", t, func() {
		a := 2
		b := 4
		Convey("When the numbers are added together", func() {
			c := a + b
			Convey("Then the result should also be an even number", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
