package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
*
题目描述
通常使用多行的节点、 父节点只 表示一棵树，比如
西安 陕西 陕西 中国 江西 中国 中国 亚洲 泰国 亚洲
输入一个节点之后，请打印出来树中他的所有下层节点
输入描述
第一行输入行数
接着是多行数据，每行以空格区分节点和父节点
最后是查询节点
树中的节点是唯一的，不会出现两个节点，是同一个名字
输出描述
输出查询节点的所有下层节点。以字典序排序
例如
5
b a
c a
d c
e c
f d
c

c是查询节点

输出下层节点，按字典序
d
e
f
*/

/*
*
思路
采用dfs或bfs都可以
*/

func main() {
	// 建树先考虑怎么表示树, 这里用哈希，key是父节点，value是他的所有子节点
	tree := make(map[string][]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		tmnp := strings.Fields(scanner.Text())
		if tree[tmnp[1]] == nil {
			tree[tmnp[1]] = make([]string, 0)
		}
		tree[tmnp[1]] = append(tree[tmnp[1]], tmnp[0])
	}
	fmt.Println(tree)
	scanner.Scan()
	search := scanner.Text()

	// 初始化结果
	ans := []string{}

	dfs(&ans, search, tree)
	//bfs(&ans, search, tree)

	sort.Strings(ans)
	fmt.Println(ans)
}

func dfs(ans *[]string, search string, tree map[string][]string) {
	*ans = append(*ans, search)
	for _, v := range tree[search] {
		dfs(ans, v, tree)
	}
}

// 使用bfs也可以
// bfs函数
func bfs(ans *[]string, target string, tree map[string][]string) {
	// 初始化队列，包含节点target
	q := []string{target}
	visited := make(map[string]bool) // 用于避免重复访问节点
	visited[target] = true

	// 进行BFS
	for len(q) > 0 {
		// 弹出队头元素node
		node := q[0]
		q = q[1:]

		// 考虑node的所有子节点
		for _, child := range tree[node] {
			if !visited[child] {
				// 将子节点加入q和ans中
				q = append(q, child)
				*ans = append(*ans, child)
				visited[child] = true
			}
		}
	}
}
