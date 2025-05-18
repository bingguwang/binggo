package main

import "fmt"

/*
*
数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
每当爬上一个阶梯都要花费对应的体力值，一旦支付了相应的体力值，就可以选择向上爬一个阶梯或者爬两个阶梯。
请找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

示例 1：
输入：cost = [10, 15, 20]
输出：15
解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。

	示例 2：

输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出：6
解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。
*/
func main() {
	cost := []int{10, 15, 20}
	stairs := minCostClimbingStairs(cost)
	fmt.Println(stairs)
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 1 {
		return cost[0]
	}
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	// dp[i]表示到达第i层阶梯的时候，最低的花费
	//dp[i] = min(dp[i-2]+cost[i], dp[i-1]+cost[i])
	for j := 2; j < len(cost)+1; j++ {
		dp[j] = min(dp[j-2]+cost[j-2], dp[j-1]+cost[j-1])
	}
	fmt.Println(dp)
	return dp[len(cost)]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
