package main

import (
	"fmt"
	"sort"
)

// maxSpending 函数：计算在不超过预算的情况下可以花费的最大金额
func maxSpending(prices []int, budget int) int {
	// 检查输入数组是否为空或元素不足，如果是，返回 -1
	if len(prices) < 3 {
		return -1 // 处理空数组或不足三个元素的情况
	}

	// 对价格数组进行排序，以便后续使用双指针技术
	sort.Ints(prices) // 按价格从低到高排序
	n := len(prices)
	maxAmount := -1 // 初始化最大花费金额为 -1

	// 遍历所有可能的第一个商品的索引
	for i := 0; i < n-2; i++ {
		left := i + 1  // 第二个商品的起始索引（左指针）
		right := n - 1 // 第三个商品的起始索引（右指针）

		// 使用双指针遍历所有可能的第二个和第三个商品的组合
		for left < right {
			total := prices[i] + prices[left] + prices[right]
			// 如果当前组合的总价格不超过预算，则更新最大花费金额，并尝试增加总价
			if total <= budget {
				if total > maxAmount {
					maxAmount = total
				}
				left++ // 增加总价
			} else {
				right-- // 减小总价
			}
		}
	}

	// 返回可以花费的最大金额
	return maxAmount
}

func main() {

	prices := []int{23, 26, 36, 27}
	var budget = 78

	// 调用 maxSpending 函数计算在给定预算下最大花费金额，并输出结果
	result := maxSpending(prices, budget)
	fmt.Println("最大花费金额:", result)
}
