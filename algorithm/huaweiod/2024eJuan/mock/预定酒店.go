package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	k, x := 4, 6
	handleBookHotel(nums, k, x)
}

func handleBookHotel(nums []int, k, x int) {
	// 自定义排序规则
	sort.Slice(nums, func(i, j int) bool {
		diff1 := math.Abs(float64(nums[i] - x))
		diff2 := math.Abs(float64(nums[j] - x))
		if diff1 != diff2 {
			return diff1 < diff2
		}
		// 与目标值差距相同时选值较小的
		return nums[i] < nums[j]
	})
	fmt.Println("nums--", nums)
	// 取前 k 个元素作为答案列表
	ans := nums[:k]

	// 对 ans 按照从小到大排序
	sort.Ints(ans)

	// 输出结果
	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i])
		if i != len(ans)-1 {
			fmt.Print(" ")
		}
	}
}
