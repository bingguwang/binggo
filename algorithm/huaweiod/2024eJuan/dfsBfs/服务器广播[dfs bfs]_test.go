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
服务器连接方式包括直接相连，间接连接。A 和 B 直接连接，B 和 C 直接连接，则 A 和 C 间接连接。直接连接和间接连接都可以发送广播。
给出一个 N * N 数组，代表 N 个服务器，matrix[i][j] == 1，则代表 i 和 j 直接连接；不等于 1 时，代表 i 和 j 不直接连接。
matrix[i][i] == 1，即自己和自己直接连接。matrix[i][j] == matrix[j][i]。
计算初始需要给几台服务器广播，才可以使每个服务器都收到广播。
输入描述	输入描述输入为 N 行，每行有 N 个数字，为 0 或 1，由空格分隔，构成 N * N 的数组，N 的范围为 1 <= N <= 50。
输出描述	输出一个数字，为需要广播的服务器数量。
------------------------------------------------------
示例1
输入
1 0 0
0 1 0
0 0 1
输出	3
说明	3 台服务器相互不连接，所以需要分别广播这 3 台服务器。

示例2
输入
1 1
1 1
输出	1
说明	2 台服务器相互连接，所以只需要广播其中一台服务器。

思路
其实就是 找出这个无向图有几个连通分量。先把 图表示出来
找连通分量可以用dfs
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// 输入第一行并解析为整数切片
	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	n := len(firstLine)
	isConnected := make([][]int, n)

	// 将第一行加入isConnected
	for _, v := range firstLine {
		val, _ := strconv.Atoi(v)
		isConnected[0] = append(isConnected[0], val)
	}

	// 输入剩余的n-1行
	for i := 1; i < n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		for _, v := range line {
			val, _ := strconv.Atoi(v)
			isConnected[i] = append(isConnected[i], val)
		}
	}
	fmt.Println("isConnected--	")
	fmt.Println(isConnected)
	var ans int
	visited := make([]bool, len(isConnected))
	for city, isVisited := range visited {
		if !isVisited { // 未访问时
			ans++
			dfsLianTong(city, isConnected, visited) //DFS函数的作用是遍历了与该服务器直接和间接相连的所有服务器,即一个连通分量
		}
	}
	fmt.Println(ans)
}

func dfsLianTong(from int, isConnected [][]int, visited []bool) {
	// 对于传入的服务器x，将其标记为已检查过
	visited[from] = true
	for to, isConnect := range isConnected[from] {
		if isConnect == 1 && !visited[to] {
			// 遍历所有与from相连的服务器
			dfsLianTong(to, isConnected, visited)
		}
	}
}

func bfsLianTong(from int, isConnected [][]int, visited []bool) {

}
