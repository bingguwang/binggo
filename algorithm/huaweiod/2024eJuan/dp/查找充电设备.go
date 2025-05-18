package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**

某个充电站，可提供n个充电设备，每个充电设备均有对应的輸出功率。
任意个充电设备组合的输出功率总和，均构成功率集合P的1个元素。功率集合P的最优元秦，表示最接近充电站最大输出功率p_max的元素。

输入描述

输入为3行：

第1行为充电设备个数n.

第2行为每个充电设备的输出功率。

第3行为充电站最大输出功率p_maxₒ

输出格式
输出功率集合
P 的最优元素。

*/

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s := scan.Text()
	n, _ := strconv.Atoi(s)
	scan.Scan()
	s2 := scan.Text()
	s3 := strings.Fields(s2)
	var nums []int
	for _, v := range s3 {
		t, _ := strconv.Atoi(v)
		nums = append(nums, t)
	}
	fmt.Println(n)
	fmt.Println(nums)
	scan.Scan()
	max, _ := strconv.Atoi(scan.Text())

	// 背包容量就是max
	// dp[i][j] 从前i个设备内选择，总功率是j的时候，可以凑出的最大输出功率
	dp := make([]int, max+1)
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := max; j >= nums[i]; j-- {
			// 第i个设备开启或不开启【这个是思路的关键】
			// dp[i][j] = maxx(dp[i-1][j], dp[i][j-nums[i]]+nums[i])
			dp[j] = maxx(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	fmt.Println(dp)
	for j := max; j >= 0; j-- {
		if dp[j] == j { // 可以凑出j
			fmt.Println(dp[j])
			return
		}
	}
	fmt.Println(0)

}
func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
