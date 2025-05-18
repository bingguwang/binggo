package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 和力扣的435 类似

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

	// 初始化答案变量
	ans := 0

	// 初始化【上个区间结束时间】为preEnd = -math.MaxInt64
	preEnd := -math.MaxInt64

	// 遍历所有区间的起始时间和结束时间
	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		// 如果【当前起始时间】大于等于【上次结束时间】
		// 可以选择【当前区间】进行观看，接在【上个区间】后面
		// 同时 preEnd 应该修改为【当前结束时间】
		// 作为下一个区间的【上次结束时间】
		if start >= preEnd {
			ans++
			preEnd = end
		} else if start < preEnd && preEnd <= end {
			// 如果【上次结束时间】正好落在【当前区间】内
			// 则不能选择【当前区间】进行观看，保留【上个区间】
			// 无需做任何事情
			continue
		} else if preEnd > end {
			// 如果【上次结束时间】大于【当前结束时间】
			// 则应该选择【当前区间】进行观看，而不应该选择【上个区间】
			// 故 preEnd 应该修改为【当前结束时间】
			// 作为下一个区间的【上次结束时间】
			preEnd = end
		}
	}

	// 输出结果
	fmt.Println(ans)
}
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)

	preEnd := -math.MaxInt
	var ans int
	for _, v := range intervals {
		start, end := v[0], v[1]
		// 如果【当前起始时间】大于等于【上次结束时间】
		// 可以选择【当前区间】进行观看，接在【上个区间】后面
		// 同时 preEnd 应该修改为【当前结束时间】
		// 作为下一个区间的【上次结束时间】
		if start >= preEnd {
			ans++
			preEnd = end
		} else if start < preEnd && preEnd <= end {
			// 如果【上次结束时间】正好落在【当前区间】内
			// 则不能选择【当前区间】进行观看，保留【上个区间】
			// 无需做任何事情
			continue
		} else if preEnd > end {
			// 如果【上次结束时间】大于【当前结束时间】
			// 则应该选择【当前区间】进行观看，而不应该选择【上个区间】
			// 故 preEnd 应该修改为【当前结束时间】
			// 作为下一个区间的【上次结束时间】
			preEnd = end
		}
	}
	fmt.Println(len(intervals) - ans)
	return len(intervals) - ans
}
