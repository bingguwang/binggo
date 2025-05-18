package main

import (
	"fmt"
	"math"
)

func main() {
	dfs(10, 0)
	fmt.Println(res)
}

var (
	res = math.MaxInt
)

func dfs(num, curtime int) { // 传入当前要处理的数值，以及当前花费的次数
	// 如果 num 为 1，则更新答案
	if num == 1 {
		// 完成一次分糖
		if curtime < res {
			res = curtime
		}
		return
	}

	// 如果 num 是偶数，递归处理 num/2
	if num%2 == 0 {
		dfs(num/2, curtime+1)
	} else {
		// 如果 num 是奇数，分别递归处理 (num+1)/2 和 (num-1)/2
		dfs((num+1)/2, curtime+2) // 做了两次操作，+1和/2
		dfs((num-1)/2, curtime+2) // 做了两次操-1和/2
	}
}
