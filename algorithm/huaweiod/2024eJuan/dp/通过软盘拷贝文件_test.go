package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const cap = 1474560 / 512 // 软盘的块数

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s2 := s.Text()
	n, _ := strconv.Atoi(s2)
	var nums []int
	for i := 0; i < n; i++ {
		s.Scan()
		teText := s.Text()
		t, _ := strconv.Atoi(teText)
		nums = append(nums, t)
	}
	fmt.Println(nums)
	// 文件占有的块数，就是文件的重量
	// 软盘的容量就是背包容量
	// 求最大能装的文件数
	var weigth []int
	for _, v := range nums {
		w := int(math.Ceil(float64(v) / 512))
		weigth = append(weigth, w)
	}
	fmt.Println(weigth)

	// dp[j]表示在软盘容量是j时，最多的文件数
	dp := make([]int, cap+1)
	for i := 0; i < len(nums); i++ {
		for j := cap; j >= weigth[i]; j-- {
			dp[j] = max(dp[j], dp[j-weigth[i]]+nums[i])
		}
	}
	fmt.Println(dp[cap])

}
