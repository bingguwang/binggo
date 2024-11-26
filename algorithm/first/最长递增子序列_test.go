package first

import (
	"fmt"
	"testing"
)

// 使用DP 可以计算

func TestName(t *testing.T) {
	//lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	// 最长递增子序列
	lengthOfLIS([]int{186, 186, 150, 200, 160, 130, 197, 200})
	// 最长递减子序列
	longestDecreasingSubsequence([]int{186, 186, 150, 200, 160, 130, 197, 200})
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ { // dp【i】表示以 nums 中第 i  个元素结尾的最长递增子序列的长度，
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			if nums[j] < nums[i] {
				dp[i] = getmax(dp[j]+1, dp[i])
			}
		}
	}
	maxval := -1
	for i := 0; i < len(dp); i++ {
		if maxval < dp[i] {
			maxval = dp[i]
		}
	}
	fmt.Println(dp)
	return maxval
}

// 最长递减子序列，有点区别，因为是dp[i]表示的是以i为起始了
func longestDecreasingSubsequence(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				dp[i] = getmax(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = getmax(res, dp[i])
	}
	fmt.Println(dp)
	return res
}

func getmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
