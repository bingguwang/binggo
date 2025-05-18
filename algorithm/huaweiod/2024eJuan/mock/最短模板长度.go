package main

import (
	"fmt"
	"sort"
)

func main() {
	handlezuid([]int{4, 5, 3, 5, 5}, 5, 3)
}

func handlezuid(heights []int, n int, rest int) {
	// 对heights数组进行从小到大排序
	sort.Ints(heights)
	// 在末尾添加一个极大值
	heights = append(heights, 2000000)
	fmt.Println(heights)
	var ans int

	for i := 0; i < n; i++ {
		// 够不够补偿所有短板到 后面一个长度的
		if rest > (heights[i+1]-heights[i])*(i+1) { // 够
			rest -= (heights[i+1] - heights[i]) * (i + 1)
		} else { // 不够则找到结果了
			ans = heights[i] + rest/(i+1) // 不够到后面一个长度，但是还是可以加几层的
			break
		}
	}

	fmt.Println(ans)
}
