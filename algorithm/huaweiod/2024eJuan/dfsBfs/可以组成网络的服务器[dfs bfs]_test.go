package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**

题目描述
在一个机房中，服务器的位置标识在 n*”的整数矩阵网格中，1表示单元格上有服务器，0 表示没有。如果两台服务器位于同一行或者同一列中紧邻的位置，则认为它们之间可以组成一个局域网
。请你统计机房中最大的局域网包含的服务器个数。

输入描述
第一行输入两个正整数，n和m，0<n,m<= 100
之后为 n*m 的二维数组，代表服务器信息
输出描述
最大局域网包含的服务器个数。
*/
// 这个和  algorithm/Backtracking/LeetCode 200. 岛屿数量_test.go 的岛屿数量那道题基本一致
var (
	rows int
	cols int
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	rows, _ = strconv.Atoi(fields[0])
	cols, _ = strconv.Atoi(fields[1])
	var arr = make([][]int, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]int, cols)
		scanner.Scan()
		tmps := strings.Fields(scanner.Text())
		for j := 0; j < cols; j++ {
			arr[i][j], _ = strconv.Atoi(tmps[j])
		}
	}
	fmt.Println(arr)
	visited := make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, cols)
	}

	var ans int
	// 开始
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if arr[i][j] == 1 && !visited[i][j] {
				var cur int
				dfsF(i, j, arr, visited, &cur)
				if cur > ans {
					ans = cur
				}
			}
		}
	}
	fmt.Println(ans)
}

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func dfsF(x, y int, arr [][]int, visited [][]bool, cur *int) {
	visited[x][y] = true //标记访问过了
	*cur++               // 当前网络机器数加一
	for _, direction := range directions {
		nx, ny := x+direction[0], y+direction[1]
		if nx >= 0 && nx < rows && ny < cols && ny >= 0 && arr[nx][ny] == 1 && !visited[nx][ny] { // 未访问，走得通
			dfsF(nx, ny, arr, visited, cur)
		}
	}
}
