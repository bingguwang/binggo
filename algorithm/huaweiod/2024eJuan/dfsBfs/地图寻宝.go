package main

import "fmt"

/*
*
小华按照地图去寻宝，地图上被划分成m行和n列的方格，横纵坐标范围分别是[0,n-1]和[0,m-1]。
在横坐标和纵坐标的数位之和不大于k的方格中存在黄金(每个方格中仅存在一克黄金)，
但横坐标和纵坐标之和大于k的方格存在危险不可进入。小华从入口(0,0)进入，任何时候只能向左，右，上，下四个方向移动一格。

请问小华最多能获得多少克黄金?
*/
func main() {
	n, m, k = 40, 40, 18
	handle(n, m, k)
}

func handle(n, m, k int) {
	used := make([][]bool, n)
	for i := range used {
		used[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !used[i][j] && compute(i)+compute(j) <= k {
				dfs(i, j, used)
			}
		}
	}
	fmt.Println(count)
}

var (
	n, m, k int
	count   int
)

func dfs(i, j int, used [][]bool) {
	if i < 0 || j < 0 || i >= n || j >= m || used[i][j] || compute(i)+compute(j) > k {
		return
	}
	used[i][j] = true
	if compute(i)+compute(j) <= k {
		count++
	}
	dfs(i-1, j, used)
	dfs(i+1, j, used)
	dfs(i, j+1, used)
	dfs(i, j-1, used)
}

func compute(a int) int {
	var res int
	for a != 0 {
		res += a % 10
		a = a / 10
	}
	return res
}
