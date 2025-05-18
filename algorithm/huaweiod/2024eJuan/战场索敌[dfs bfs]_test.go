package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
*
这个和 algorithm/Backtracking/LeetCode 200. 岛屿数量_test.go 的岛屿数量那道题差不多
*/
var rows, cols int

func main() {
	// 准备工作
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	rows, _ = strconv.Atoi(fields[0])
	cols, _ = strconv.Atoi(fields[1])
	target, _ := strconv.Atoi(fields[2])
	arr := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		scanner.Scan()
		arr[i] = []rune(scanner.Text())
		fmt.Println(arr[i])
	}
	// 逻辑

	// 因为连通域里的元素，其实在计算上得到的结果是一样的，所以使用标记数组
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	var ans int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && (arr[i][j] == '.' || arr[i][j] == 'E') {
				curNum := 0
				dfsFind(i, j, arr, visited, &curNum)
				if target > curNum {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}

var DIRECTIONS = [4][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func dfsFind(x, y int, arr [][]rune, visited [][]bool, cursum *int) {

	// 将点(x, y)标记为已检查过
	visited[x][y] = true
	// 如果当前位置是敌人，则更新当前连通块的敌人数目curNum
	if arr[x][y] == 'E' {
		*cursum++
	}
	for _, dir := range DIRECTIONS {
		nxtX, nxtY := x+dir[0], y+dir[1]
		// 若下一个点继续进行dfs，应该满足以下三个条件：
		// 1. 没有越界
		// 2. 在grid中值为"."或"E"
		// 3. 尚未被检查过
		if nxtX >= 0 && nxtX < len(arr) && nxtY >= 0 && nxtY < len(arr[0]) {
			if (arr[nxtX][nxtY] == '.' || arr[nxtX][nxtY] == 'E') && !visited[nxtX][nxtY] {
				// 可以进行dfs
				dfsFind(nxtX, nxtY, arr, visited, cursum)
			}
		}
	}
}
