package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 初始化输入读取器
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// 输入 n 和 m
	parts := strings.Split(input, " ")
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	// 输入剩余 n 张牌
	numsInput, _ := reader.ReadString('\n')
	numsInput = strings.TrimSpace(numsInput)
	numsStr := strings.Split(numsInput, " ")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(numsStr[i])
	}

	// 根据 isFind 的结果，输出数字 0 或 1
	fmt.Println(presumHandle(nums, m))
}

// 关键在于知道前缀和的公式
// sum(nums[i:j]) = pre_sum_lst[j] - pre_sum_lst[i]
// 于是依照题目就有 sum(nums[i:j]) %m = (pre_sum_lst[j] - pre_sum_lst[i])%m=0
// 就是  pre_sum_lst[j]%m =  pre_sum_lst[i] %m
// 所以这里吧前缀和%m存在哈希里，找到存在的就可以

func presumHandle(nums []int, m int) bool {
	// 设置一个集合，用来储存所有前缀和对 m 的求余结果
	preSumSet := make(map[int]bool)
	// 前缀和 0 始终可以取得到，即不选取任何一个数字，0 % m = 0，在集合中储存 0
	preSumSet[0] = true

	// 初始化前缀和为 0
	preSum := 0
	// 初始化标志，表示是否找到一段连续的数组可以整除
	isFind := false

	// 遍历 nums 数组
	for _, num := range nums {
		// 前缀和加上 num
		preSum += num
		fmt.Println("前缀和是:", preSum)
		fmt.Println("前缀和%7是:", preSum%m)
		// 如果 preSum % m 的结果位于 preSumSet 中
		if _, exists := preSumSet[preSum%m]; exists {
			isFind = true
			break
		}
		// 如果没有进入上述 if，则需要把 preSum % m 的结果储存入集合 preSumSet 中
		preSumSet[preSum%m] = true

		fmt.Println(preSumSet)
	}
	return isFind
}
