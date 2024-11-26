package od

import (
	"fmt"
	"testing"
)

// 其实就是最长递增子序列的变形

func TestName22(t *testing.T) {
	var n int
	fmt.Scan(&n)
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ints[i])
	}
	//lengthOfLIS([]int{186, 186, 150, 200, 160, 130, 197, 200})
	ints = []int{186, 186, 150, 200, 160, 130, 197, 200}
	l1 := lengthOfLIS(ints)
	l2 := longestDecreasingSubsequence(ints)
	var rt = -1
	for i, v := range l1 {
		if rt < l2[i]+v {
			rt = l2[i] + v
		}
	}
	fmt.Println(n - rt + 1)
}

func lengthOfLIS(nums []int) []int {
	dp := make([]int, len(nums))
	//dp[i]表示到第i个数为止，最长递增子序列的长度
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			if nums[j] < nums[i] {
				dp[i] = getMax2(dp[j]+1, dp[i])
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
	return dp
}

func longestDecreasingSubsequence(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] { // 满足递减序
				dp[i] = getMax2(dp[i], dp[j]+1)
			}
		}
	}
	res := 0 // 最大值
	for i := 0; i < n; i++ {
		res = getMax2(res, dp[i])
	}
	fmt.Println(dp)
	return dp
}

func getMax2(i, j int) int {
	if i > j {
		return i
	}
	return j
}
