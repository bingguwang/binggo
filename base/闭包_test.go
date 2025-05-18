package base

import (
	"fmt"
	"testing"
)

func createCounter(start int) func(int) int {
	current := start
	fmt.Println("current：", current)
	return func(inc int) int {
		current += inc
		return current
	}
}
func TestNamessw(t *testing.T) {
	counter := createCounter(10)
	fmt.Println(counter(5)) // 输出: 15
	fmt.Println(counter(3)) // 输出: 18
}
