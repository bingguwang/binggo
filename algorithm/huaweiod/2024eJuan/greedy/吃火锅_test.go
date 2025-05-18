package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i] = make([]int, 2)
		f := strings.Fields(scanner.Text())
		nums[i][0], _ = strconv.Atoi(f[0])
		nums[i][1], _ = strconv.Atoi(f[1])
		nums[i][1] += nums[i][0]
	}
	fmt.Println(nums)
	fmt.Println(m)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][1] < nums[j][1]
	})
	fmt.Println(nums)

	preend := -1
	var ans int
	for _, num := range nums {
		// 如果当前某个菜刚好合适，而且距离上次吃菜的时间超过了手速 m
		// 那么直接选择这个菜，更新答案变量，且修改 preTime 变量
		// 表示对于下一次吃菜，当前的 t 成为了其上一次吃菜时间 preTime
		if preend+m <= num[1] {
			preend = num[1]
			ans++
		}
	}

	fmt.Println(ans)
}
