package main

import (
	"fmt"
)

func main() {
	fmt.Println(shuixianhuashu(3, 0))
	fmt.Println(shuixianhuashu(9, 1))
}

func shuixianhuashu(n, m int) int {
	if n < 3 || n > 7 {
		return -1
	}
	// n位数
	/**
	  n=3 从100, 101, 102, .... 999 开始找
	  n=4 从1000, 1001, 102, .... 9999 开始找
	*/
	start := pow(10, n-1)
	end := pow(10, n)
	fmt.Println("start:", start)
	var count = -1
	for i := start; i < end; i++ {
		if check(i, n) {
			count++
			if count == m {
				return i
			}
		}
	}
	return -1
}

// 检查某个数a是否是水仙花数，指数是n
func check(a int, n int) bool {
	// 153
	var tmp int
	mo := a
	for a != 0 {
		c := a % 10       // 3
		tmp += pow2(c, n) // +3^3
		a = a / 10        // 15, 1
	}
	return mo == tmp
}
func pow2(x, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res = x * res
	}
	return res
}
