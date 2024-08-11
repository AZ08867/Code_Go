package series

import "fmt"

func GetFibonacciSeries(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

func Square(n int) int {
	return n * n
}

func init() {
	fmt.Println("init")
}
func init() {
	fmt.Println("init2")
}
func init() {
	fmt.Println("init3")
}
