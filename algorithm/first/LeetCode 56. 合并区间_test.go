package first

import (
    "fmt"
    "sort"
    "testing"
)

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/
func TestFAJSDA(t *testing.T) {
    a := [][]int{
        {1, 3},
        {2, 6},
        {8, 10},
        {15, 18},
    }
    number := merge(a)
    fmt.Println(number)
}
func merge(intervals [][]int) [][]int {
    // 按照起始位置从小到大排序
    sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
    fmt.Println(intervals)
    // 初始化答案列表，将第一个区间加入其中
    ans := [][]int{intervals[0]}
    // 遍历每个区间
    for i := 1; i < len(intervals); i++ {
        // 如果当前区间和前一个区间重叠，则合并它们
        if intervals[i][0] <= ans[len(ans)-1][1] {
            ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
        } else {
            // 否则将前一个区间加入答案并更新当前区间
            ans = append(ans, intervals[i])
        }
    }
    return ans
}
