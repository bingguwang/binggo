package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1,-5,-6,4,3,6,-2
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	split := strings.Split(scanner.Text(), ",")
	nums := make([]int, len(split))
	for _, v := range split {
		atoi, _ := strconv.Atoi(v)
		nums = append(nums, atoi)
	}
	n := len(nums)
	// fi怎么设置
	// fi表示第i轮结束时的最大分数
	var f = make([]int, len(nums))

	// 状态方程
	// f(i) = max(f(i-1)+score[i] ,f(i-3)  ) i>=3
	// f(i) = max(f(i-1)+score[i] ,0  ) i<3

	// 怎么初始化
	if nums[0] > 0 {
		f[0] = nums[0]
	} else {
		f[0] = 0
	}

	for i := 1; i < len(nums); i++ {
		if i >= 3 {
			f[i] = maxval(f[i-1]+nums[i], f[i-3])
		} else {
			f[i] = maxval(f[i-1]+nums[i], 0)
		}
	}
	fmt.Println(f)
	fmt.Println(f[n-1])
}

func maxval(a, b int) int {
	if a > b {
		return a
	}
	return b
}
