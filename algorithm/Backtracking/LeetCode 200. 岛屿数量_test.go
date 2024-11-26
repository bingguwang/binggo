package Backtracking

import (
    "fmt"
)

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/
func main() {
    a := [][]byte{
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'1', '1', '0', '0', '0'},
        {'0', '0', '0', '0', '0'},
    }
    number := numIslands(a)
    fmt.Println(number)
}
func numIslands(grid [][]byte) int {
    n := len(grid)
    m := len(grid[0])
    visited := make([][]bool, len(grid))
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, len(grid[0]))
    }
    var dfs func(i, j int)
    dfs = func(i, j int) {
        // 越界?
        if i >= n || j >= m || i < 0 || j < 0 || visited[i][j] || grid[i][j] == '0' {
            return
        }
        visited[i][j] = true
        // 把所有能走的都走过
        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j-1)
        dfs(i, j+1)
    }
    var res int
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                dfs(i, j)
                res++
            }
        }
    }
    return res
}