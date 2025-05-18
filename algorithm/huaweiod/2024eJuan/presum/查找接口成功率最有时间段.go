package main

import (
	"fmt"
	"strings"
)

func main() {
	solve(1, []int{0, 1, 2, 3, 4})
}

func handle(nums []int, minavg int) [][]int {
	presum := make([]int, len(nums)+1)
	presum[0] = 0

	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	fmt.Println(presum)
	var res [][]int
	// 得到前缀和后，sum(nums[i:j]) =pre[j]-pre[i]
	// sum(nums[i:j]) / 3 <= minavg
	// (pre[j]-pre[i])/3 <= minavg
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (presum[j]-presum[i])/(j-i) <= minavg {
				res = append(res, []int{i, j - 1})
			}
		}
	}
	fmt.Println(res)
	return res
}

// 构建解决问题的函数
func solve(minAverageLost int, nums []int) string {
	// 数据长度
	n := len(nums)

	// 构建前缀和数组，注意首位需要填充一个0，表示不选取任何数字的前缀和
	preSum := make([]int, n+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	fmt.Println(preSum)
	// 构建答案数组
	var ans []string

	// 逆序遍历区间的长度 winlen ，贪心地优先考虑尽可能大的区间
	// // 窗口长度 最大为n, 最小是1, 因为要找出最大的所以倒序找
	for winlen := n; winlen > 0; winlen-- {

		// 遍历区间的起始位置 i，其范围为 [0, n-length+1)
		for left := 0; left <= n-winlen; left++ {
			// 对于每一个区间的起始位置 i，我们都需要考虑长度为 l 的区间 [i:i+l] 的区间和
			//得到前缀和后，sum(nums[i:j]) =pre[j]-pre[i]
			// 使用前缀和计算区间和 intervalSum
			right := left + winlen
			intervalSum := preSum[right] - preSum[left]

			// 如果区间和小于等于阈值，则这个区间是满足题意的区间，将其加入 ans 中
			if intervalSum <= minAverageLost*winlen {
				// 储存的区间是左闭右闭区间，故右边界应该为 i+length-1
				ans = append(ans, fmt.Sprintf("%d-%d", left, right-1))
			}
		}

		// 在考虑大小为 length 的区间之后，如果 ans 中有值
		// 则说明找到了最长的满足题意的区间，将 ans 合并后返回输出
		if len(ans) > 0 { // 长度由大到小遍历，找到就不用遍历小长度的了
			fmt.Println(ans)
			return strings.Join(ans, " ")
		}
	}

	// 如果退出循环后，没有返回任何的一个 ans，则说明找不到任意一个区间满足题意
	// 此时应该返回 "NULL" 输出
	return "NULL"
}
