package main

import (
	"fmt"
	"sort"
)

func main() {
	m := 3                           // m个月内完成
	requirement := []int{3, 5, 3, 4} // 需求的工作量
	sort.Ints(requirement)
	sum := 0
	for i := range requirement {
		sum += requirement[i]
	}

	// 每个月最多只开发2个需求
	// 求每个月的人力
	// k的范围先确定, 最大是  , 最小是
	// 二分查找
	left, right := requirement[len(requirement)-1], sum+1
	for left < right {
		mid := (left + right) / 2
		if handle(requirement, mid) > m { // 速度不够，人力加大
			left = mid + 1 //(小的一方进行相加)
		} else {
			right = mid
		}
	}
	// 退出二分查找时，k = right = left
	fmt.Println(left)
	fmt.Println(right)
}

// 假设说每个月的人力是k
func handle(requirement []int, k int) int {
	// 需要在m个月里完成工作, 挑一个或2个
	var res int
	left, right := 0, len(requirement)-1
	for left <= right {
		// 大值加小值
		// 超过每个月的人力了 ,只能把大的工作先做了
		if requirement[left]+requirement[right] > k {
			right--
		} else { // 没超过限制，两个工作都可以放到一个月里做
			left++
			right--
		}
		res++
	}
	return res
}
