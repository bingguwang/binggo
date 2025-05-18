package main

import (
	"sort"
)

func handleZhengshuDuiMin(arr1 []int, arr2 []int, k int) int {
	var ans []int
	// 暴力求解
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			ans = append(ans, arr1[i]+arr2[j])
		}
	}
	sort.Ints(ans)
	var sum int
	for i := 0; i < k; i++ {
		sum += ans[i]
	}
	return sum
}
