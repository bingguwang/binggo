package main

/*
*
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

一开始你在下标 0 处。每一步，你最多可以往前跳 k 步，但你不能跳出数组的边界。也就是说，你可以从下标 i 跳到 [i + 1， min(n - 1, i + k)] 包含 两个端点的任意位置。

你的目标是到达数组最后一个位置（下标为 n - 1 ），你的 得分 为经过的所有数字之和。

请你返回你能得到的 最大得分 。
*/
func maxResult(nums []int, k int) int {
	n := len(nums)

	// 初始化 dp 数组
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1 << 31 // 设置为一个极小值
	}
	// 使用切片模拟双端队列，存储索引
	queue := []int{}
	// 动态规划计算
	for i := 0; i < n; i++ {
		// 如果队列为空，直接初始化 dp[i]
		if len(queue) == 0 {
			dp[i] = nums[i]
		} else {
			// 移除超出窗口范围的元素
			for len(queue) > 0 && queue[0] < i-k {
				queue = queue[1:]
			}
			// 更新 dp[i]
			dp[i] = dp[queue[0]] + nums[i]
		}
		// 移除队列中比当前值小的元素（保持单调递减）
		for len(queue) > 0 && dp[queue[len(queue)-1]] < dp[i] {
			queue = queue[:len(queue)-1]
		}
		// 将当前位置加入队列
		queue = append(queue, i)
	}

	// 返回最后一个位置的最大得分
	return dp[n-1]
}
