package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
*
在通信系统中，一个常见的问题是对用户进行不同策略的调度，会得到不同的系统消耗和 性能Q假设当前有n个待串行调度用户，每个用户可以使用A/B/C三种不同的调度策略，不同的策略会消耗不同的系统资源。请你根据如下规则进行用户调度，并返回总的消耗资源数。
规则:
相邻的用户不能使用相同的调度策略，例如，第1个用户使用了A策略，则第2个用户只能使用B或者c策略。对单个用户而言，不同的调度策略对系统资源的消耗可以归一化后抽象为数值。例如，某用户分别使用 B/ 策略的系统消耗分别为15/8/17 。

每个用户依次选择当前所能选择的对系统资源消耗最少的策略(局部最优)，如果有多个满足要求的策略，选最后一个

。
输入描述
第一行表示用户个数 n
接下来每一行表示一个用户分别使用三个策略的系统消耗resA rese resc
输出描述
最优策略组合下的总的系统资源消耗数
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取用户个数
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([][]int, n)

	// 读取每个用户的资源消耗
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		nums[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			nums[i][j], _ = strconv.Atoi(input[j])
		}
	}

	// 准备工作完成
	fmt.Println("nums---", nums)

	// 初始化第一次选择
	curChoice := find(nums[0], -1)
	ans := nums[0][curChoice]
	preChoice := curChoice

	// 遍历剩余用户的选择情况
	for i := 1; i < n; i++ {
		curChoice = find(nums[i], preChoice)
		ans += nums[i][curChoice]
		preChoice = curChoice
	}

	fmt.Println(ans)
}

// 查找当前选择最小的资源消耗
func find(lst []int, preChoice int) int {
	minCost := math.MaxInt64 // 初始化为无穷大
	curChoice := -1

	// 逆序遍历寻找最小花费，同时满足不能与上次选择相同
	for i := 2; i >= 0; i-- {
		if i == preChoice {
			continue
		}
		if lst[i] < minCost {
			minCost = lst[i]
			curChoice = i
		}
	}
	return curChoice
}
