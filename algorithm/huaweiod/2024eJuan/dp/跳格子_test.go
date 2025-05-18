package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
*
小明和朋友们一起玩跳格子游戏，每个格子上有特定的分数。
比如，score[]=[1,-1,-6,7,-17,7]，从起点score[8]开始，每次最大跳的步长为k，请你返回小明跳到终点score[n-1]时，能得到的最大得分。
注:
。格子的总长度和步长的区间在[1，100000]
。每个格子的分数在[-10808，10000]区间中;
输入描述
第一行输入总的格子数量n
第二行输入每个格子的分数 score[]
第三行输入最大跳的步长 k
输出描述
输出最大得分数

思路
动态规划的思路是：三部曲
f[i]
状态方程
初始化状态方程
*/
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	n, _ := strconv.Atoi(scanner.Text())
//	scanner.Scan()
//	fields := strings.Fields(scanner.Text())
//	input := []int{}
//	for _, v := range fields {
//		atoi, _ := strconv.Atoi(v)
//		input = append(input, atoi)
//	}
//	scanner.Scan()
//	k, _ := strconv.Atoi(scanner.Text())
//	var f = make([]int, n)
//
//	// 初始化
//	for i := 0; i < len(f); i++ {
//		f[i] = -1e9
//	}
//	f[0] = input[0]
//
//	for i := 1; i < n; i++ {
//		for j := 1; j <= k; j++ {
//			// 转移方程如下
//			//f[i] = max( f[i-1], f[i-2], f[i-3]...f[i-k]) + input[i]
//			if i-j >= 0 {
//				f[i] = max(f[i], f[i-j]+input[i])
//			}
//		}
//	}
//
//	fmt.Println(f)
//	fmt.Println(f[n-1])
//}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
上面的会有超时的问题

使用单调队列会比较合适

*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	input := []int{}
	for _, v := range fields {
		atoi, _ := strconv.Atoi(v)
		input = append(input, atoi)
	}
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	var dp = make([]int, n)

	// 初始化
	for i := 0; i < len(dp); i++ {
		dp[i] = -1 << 9
	}
	dp[0] = input[0]

	// 队列里存下标
	var q = []int{0}
	for i := 1; i < len(input); {
		// i不在窗口内，队头出队，窗口移动
		if q[0]+k < i {
			q = q[1:]
		}
		//状态转移
		// dp[i] = max(dp[i-1]+1, dp[i-2]+2,...dp[i-k]+k)
		// 其实就是在dp[i-k]到dp[i-1]里面找到最大值，然后加上input[i]
		dp[i] = dp[q[0]] + input[i]

		// 比队尾元素大，队尾出队。保证队首始终是最大位置的索引
		for len(q) >= 1 && dp[i] >= q[len(q)-1] {
			// 出队
			q = q[:len(q)-1]
		}
		q = append(q, i) //入队
		i++              // 窗口向后移动一位
	}
	fmt.Println(q)
	fmt.Println(dp)
}
