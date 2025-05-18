package main

import (
	"fmt"
)

func main() {
	//handle([]int{123, 124, 125, 121, 119, 122, 126, 123}, 8)
	dandiaozhanhandle([]int{123, 124, 125, 121, 119, 122, 126, 123}, 8)
	//handle([]int{100, 95}, 2)
}

// 暴力法
func handle(nums []int, n int) []int {
	var res []int
	// stack := make([]int, 0) // 单调栈，存下标

	for i := 0; i < n; i++ {
		cur := nums[i]
		var flag bool
		for j := i; j < n; j++ {
			if cur < nums[j] {
				res = append(res, j)
				flag = true
				break
			}
		}
		if !flag {
			res = append(res, 0)
		}
	}
	fmt.Println(res)
	return res
}

func dandiaozhanhandle(height []int, n int) {
	// 构建一个单调栈，用来存放不同小朋友的身高的索引
	// 栈中储存的索引所对应在 height 中的元素大小，从栈底至栈顶单调递减
	stack := []int{}

	// 构建列表 ans，用来保存输出结果
	ans := make([]int, n)

	// 从头开始遍历每一个小朋友的身高
	for i := 0; i < n; i++ {
		// 第 i 个小朋友的身高 h，需要不断地与栈顶元素比较
		// 如果栈顶元素存在并且 h 【大于】 栈顶元素 height[stack[len(stack)-1]]
		// 意味着栈顶元素找到了右边最近的比他更高的身高 h
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// 首先获取栈顶元素的值，也就是上一个比 h 小的身高的索引值
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 弹出栈顶元素

			// i 即为 preIndex 这个索引所对应的，下一个最近身高
			ans[preIndex] = i
		}

		// 再把当前小朋友身高的下标 i 存放到栈中
		// 注意：所储存的是下标 i，而不是身高 h
		stack = append(stack, i)
	}
	fmt.Println(ans)

}
