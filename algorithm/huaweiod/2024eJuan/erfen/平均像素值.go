package main

import (
	"fmt"
)

// 计算新数组的平均值
func calNewAverage(nums []int, k, n int) float64 {
	// 新数组的和，初始化为0
	newSum := 0

	// 遍历原数组中的所有数字 num
	for _, num := range nums {
		// 加 k 得到新数组 newNum
		newNum := num + k

		// 若新数字小于 0，则修改为 0
		if newNum < 0 {
			newNum = 0
		}

		// 若新数字大于 255，则修改为 255
		if newNum > 255 {
			newNum = 255
		}

		// 将新数字加入 newSum 中
		newSum += newNum
	}

	// 返回新数组的和的平均值
	return float64(newSum) / float64(n)
}

func main() {
	nums := []int{129, 130, 129, 130}

	// 计算原数组的长度 n
	n := len(nums)

	// k 的左闭右开区间，right 最大值为 255，闭区间取值 256
	left, right := -255, 256

	// 二分查找，计算第一个使得新数组平均值小于 128 的 k
	for left < right {
		mid := left + (right-left)/2

		// 若计算结果小于 128，说明整体平均值还可以更大，left 右移
		if calNewAverage(nums, mid, n) < 128 {
			left = mid + 1
		} else { // 若计算结果不小于 128，说明整体平均值需要变小，right 左移
			right = mid
		}
		fmt.Println(left, " ", right)
	}

	// 退出循环后，k = left 是使得 calNewAverage(nums, k, n) 恰好大于等于（不小于）128 的值
	// left 和 left-1 都有可能是答案，看哪一个更接近 128
	avgLeftMinus1 := calNewAverage(nums, left-1, n)
	avgLeft := calNewAverage(nums, left, n)

	if 128-float64(avgLeftMinus1) <= float64(avgLeft)-128 {
		fmt.Println(left - 1)
	} else {
		fmt.Println(left)
	}
}
