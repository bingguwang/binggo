package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	arr := strings.Fields(s)
	var nums []int
	for _, v := range arr {
		t, _ := strconv.Atoi(v)
		nums = append(nums, t)
	}
	fmt.Println(nums)
	f(nums)
}

func f(nums []int) {
	n := len(nums)

	stack := make([]int, 0) // 单调栈
	// res保存结果集
	res := make([]int, n)
	copy(res, nums)

	for i := 0; i < 2*n; i++ {
		idx := i % n     // 在原数组里真正的下标
		cur := nums[idx] // 当前元素

		// 单调栈中存储的是尚未找到右侧更小元素的元素下标
		// 栈空或者栈顶对应的元素大于当前元素，也就是当前元素就是栈顶元素 右边的第一个小的值
		for len(stack) > 0 && nums[stack[len(stack)-1]] > cur {
			// 栈顶对应的元素，更新他的价值
			// 出栈
			// 循环弹出栈顶元素，直到栈为空或栈顶元素不大于 cur
			topidx := stack[len(stack)-1]
			res[topidx] = cur + nums[topidx]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, idx)
	}

	// 输出结果
	for _, v := range res {
		fmt.Print(v, " ")
	}
}
