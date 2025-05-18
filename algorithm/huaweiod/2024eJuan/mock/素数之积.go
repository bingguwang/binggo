package main

import (
	"fmt"
	"math"
)

func main() {
	n := 15
	for i := 2; i < n; i++ {
		if sushu(i) {
			other := n / i
			if sushu(other) && other*i == n {
				fmt.Println(i, " ", other)
				return
			}
		}
	}
}

// 判断某个数是否是素数
func sushu(a int) bool {
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
