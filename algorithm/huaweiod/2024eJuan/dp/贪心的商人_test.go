package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 可以贪心，可以dp

// 贪心
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	number, _ := strconv.Atoi(scanner.Text())
//	scanner.Scan()
//	days, _ := strconv.Atoi(scanner.Text())
//
//	scanner.Scan()
//	tmp := strings.Fields(scanner.Text())
//	items := make([]int, number)
//	for i, v := range tmp {
//		items[i], _ = strconv.Atoi(v)
//	}
//	goods := make([][]int, len(items))
//	for i := 0; i < number; i++ {
//		scanner.Scan()
//		f := strings.Fields(scanner.Text())
//		goods[i] = make([]int, days)
//		for j, v := range f {
//			goods[i][j], _ = strconv.Atoi(v)
//		}
//	}
//	fmt.Println(items)
//	fmt.Println(goods)
//
//	// 准备工作完成
//	fmt.Println("准备工作完成")
//
//	// 贪心，因为各个商品卖出买入是独立的
//	var ans int
//	for i := 0; i < len(goods); i++ {
//		ans += tanxin(goods[i]) * items[i]
//	}
//	fmt.Println(ans)
//}

// 求每件商品的最大利润
// 传入某件商品的每天的价格
// 返回这些天里，这件商品的最大利润
func tanxin(goods []int) int {
	maxlirun := 0
	for i := 1; i < len(goods); i++ {
		if goods[i] > goods[i-1] {
			maxlirun += goods[i] - goods[i-1]
		}
	}
	fmt.Println("maxlirun--", maxlirun)
	return maxlirun
}

/**
3
3
4 5 6
1 2 3
4 3 2
1 5 3

*/

// 计算某个特定商品能取得的最大利润的函数
func maxProfit(prices []int, days int) int {
	dp := make([][2]int, days)

	// dp[i][0]第i天不持有，dp[i][1]第i天持有
	dp[0][0] = 0          // 第一天不持有，利润为 0
	dp[0][1] = -prices[0] // 第一天持有，利润为 -prices[0]

	// 动态规划
	for i := 1; i < days; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 不持有
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 持有
	}
	fmt.Println("dp[days-1][0]:  ", dp[days-1][0])

	return dp[days-1][0] // 返回最后一天不持有股票的最大利润
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 输入商品数量
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 输入天数
	scanner.Scan()
	days, _ := strconv.Atoi(scanner.Text())

	// 输入每种商品的最大数目
	scanner.Scan()
	numbersStr := strings.Split(scanner.Text(), " ")
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i], _ = strconv.Atoi(numbersStr[i])
	}

	ans := 0
	// 循环输入每种商品的价格变化
	for i := 0; i < n; i++ {
		scanner.Scan()
		// 商品每天的价格
		pricesStr := strings.Fields(scanner.Text())
		prices := make([]int, days)
		for j := 0; j < days; j++ {
			prices[j], _ = strconv.Atoi(pricesStr[j])
		}
		// 计算单件商品的最大利润并乘以数量
		ans += maxProfit(prices, days) * numbers[i]
	}

	fmt.Println(ans)
}
