package main

import (
	"bufio"
	"fmt"
	"github.com/zeromicro/go-zero/core/mathx"
	"os"
	"strconv"
	"strings"
)

/*
*
1 2 3 4 5 6 7 8 9 10
*/
var (
	ans   int
	total int
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	nums := []int{}
	for _, v := range fields {
		atoi, _ := strconv.Atoi(v)
		nums = append(nums, atoi)
		total += atoi
	}
	avg := total / 2
	ans = total

	//dfs(nums, 0, 0, 0, avg)
	dpFun(nums, avg)
	fmt.Println(ans)
}

// 可以使用回溯
func dfs(nums []int, cursum int, chooesed int, startindex int, target int) {
	if cursum > target {
		return
	}
	if chooesed == 5 {
		ans = mathx.MinInt(total-2*cursum, ans)
		return
	}

	for i := startindex; i < len(nums); i++ {
		dfs(nums, cursum+nums[i], chooesed+1, i+1, target)
	}
}

// 也可以使用动态规划
// 0-1背包问题
func dpFun(nums []int, target int) {
	//求最小差值的公式是：ans =  total - 2*较小一组和
	//dp三部曲
	//dp[i][j]，i表示从前i个人选出若干个人， 组成评分和为j时，是否可做到
	// 我们只需要关注其中一组，于是关注和小于total/2的那组就行
	// 问题就变成在10人里选出5个人的问题
	// 对于第i个人可以选或不选两种可能
	//dp[i][j] = dp[i-1][j] // 第i个人不选，那就是从前i-1 人里选出若干个组成和为j的可能性
	//dp[i][j] = dp[i-1][j-nums[i-1]] // / 第i个人选，那就是从前i-1  人里选出若干个组成和为j-nums[i-1]的可能性，因为nums是从0开始的
	// 可以看到可以优化，和第一维的数组无关，于是可以优化为一维数组的状态转移方程 d[j]=dp[j], dp[j-nums[i]], j不超过的target
	dp := make([]bool, target+1) // 于是dp[j]
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	fmt.Println(dp)

	for j := target; j >= 0; j-- { // 从中间开始找，找到的第一个可以组成的分数就是最优解
		if dp[j] {
			ans = total - 2*j
			break
		}
	}
	fmt.Println(ans)
}
