package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 在贪心里有这道提，这道题也可以使用dp来解，也对应了力扣的第300题 最长递增子序列
func main() {
	// 读取输入
	scanner := bufio.NewScanner(os.Stdin)

	// 输入演出的数目
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// 初始化间隔列表
	intervals := make([][2]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		start, _ := strconv.Atoi(parts[0])
		during, _ := strconv.Atoi(parts[1])
		// 对于每一个结束时间都+15后再储存，方便后续进行比较
		end := start + during + 15
		intervals[i] = [2]int{start, end}
	}

	// 对intervals进行排序,按开始时间递增排序。 对于开始时间一样的，结束时间较早的排前面
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	//和下面的排序效果是一致的

	// 对intervals进行排序,按结束时间递增排序
	/*sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})*/

	fmt.Println("intervals--", intervals)

	// 转化为LIS问题，LIS问题里是根据大小来排序，这里是根据 start>=pred_end来排列
	var dp = make([]int, len(intervals)) // dp[i] 表示第i场作为最后一场演出时， 最长无重叠演出数目
	dp[0] = 1
	for i := 1; i < len(intervals); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	fmt.Println(dp)

	// 找到dp数组中的最大值
	maxWatch := 0
	for _, v := range dp {
		maxWatch = max(maxWatch, v)
	}
	fmt.Println(maxWatch)
}

//
//func main() {
//	s := bufio.NewScanner(os.Stdin)
//	s.Scan()
//	n, _ := strconv.Atoi(s.Text())
//	var showarr [][]int
//	for i := 0; i < n; i++ {
//		s.Scan()
//		var show []int
//		s2 := strings.Fields(s.Text())
//		for _, v := range s2 {
//			interval, _ := strconv.Atoi(v)
//			if len(show) == 1 {
//				interval = show[0] + interval + 15
//			}
//			show = append(show, interval)
//		}
//		showarr = append(showarr, show)
//	}
//	fmt.Println(showarr)
//	sort.Slice(showarr, func(i, j int) bool {
//		if showarr[i][0] == showarr[j][0] {
//			return showarr[i][1] < showarr[j][1]
//		}
//		return showarr[i][0] < showarr[j][0]
//	})
//
//	var ans = 0
//	dp := make([]int, len(showarr))
//	dp[0] = 1
//	// dp[i] 表示以第i场演出为最后一场时，最多能看几场
//	for i := 1; i < len(showarr); i++ {
//		dp[i] = 1
//		for j := 0; j < i; j++ {
//			if showarr[j][1] <= showarr[i][0] {
//				dp[i] = max(dp[j]+1, dp[i])
//			}
//			fmt.Println(dp)
//		}
//		if ans < dp[i] {
//			ans = dp[i]
//		}
//	}
//	fmt.Println(dp)
//	fmt.Println(ans)
//
//}
