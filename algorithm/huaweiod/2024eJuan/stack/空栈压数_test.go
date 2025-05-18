package main

import (
	"fmt"
)

/**
核心的思路就是

模拟的其实只有两个东西

不断更新
更新stack的状态
更新当前要加入值


*/
// 检查是否需要弹出若干栈顶元素并进行压缩的函数
func checkStackTop(stack []int, num int) ([]int, int, bool) {
	// 初始化栈顶元素的和为0
	topSum := 0
	// 初始化栈顶元素索引为idx，逆序遍历
	idx := len(stack) - 1

	// 进行循环，满足以下两个条件之一，则退出循环：
	// 1. 栈顶元素和大于等于num
	// 2. 逆序遍历的索引idx小于0
	for topSum < num && idx >= 0 {
		// 在循环中，
		// 栈顶元素和递增
		// 栈顶元素索引递减
		topSum += stack[idx]
		idx--
	}

	// 退出循环后，若栈顶元素和topSum恰好等于num
	if topSum == num {
		// 需要继续进行循环，故返回true
		// 需要删除之前计入栈顶元素和中的所有元素，用切片操作即可
		// num更新为原来的两倍，返回这两个参数
		return stack[:idx+1], num * 2, true
	} else {
		// 无需继续进行循环，故返回false
		// 返回原先的stack和num
		return stack, num, false
	}
}

func main() {
	// 将字符串转换为整数切片
	var nums []int = []int{6, 1, 2, 3}

	// 初始化一个空栈
	stack := []int{}

	// 从左到右，遍历每一个数字
	for _, num := range nums {
		// 设置变量 flagContinueLoop 为 true
		flagContinueLoop := true

		// 当 flagContinueLoop 为 true 时，持续循环，也就是连锁反应
		for flagContinueLoop {
			// 调用 checkStackTop 函数，更新 stack 和 num 的情况
			var updatedStack []int
			fmt.Println("栈内是:", stack)
			fmt.Println("要加入:", num)
			updatedStack, num, flagContinueLoop = checkStackTop(stack, num)
			stack = updatedStack
		}

		// 做完计算，必须把最新得到的 num 加入栈中
		stack = append(stack, num)
	}

	fmt.Print(stack)
}
