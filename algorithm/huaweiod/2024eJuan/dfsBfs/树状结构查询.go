package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// dfs函数
func dfs(node string, ans *[]string, childrenDict map[string][]string) {
	// 将node加入ans中
	*ans = append(*ans, node)
	// 遍历node的所有子节点，递归调用子节点
	for _, child := range childrenDict[node] {
		dfs(child, ans, childrenDict)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 输入行数
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(nStr)

	// 使用map存储某一个节点所有子节点构成的列表
	childrenDict := make(map[string][]string)

	// 遍历N行，输入所有的父子节点连接情况
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		child := parts[0]
		parent := parts[1]
		childrenDict[parent] = append(childrenDict[parent], child)
	}

	// 输入查询节点
	scanner.Scan()
	target := scanner.Text()

	// 初始化答案切片
	var ans []string
	// dfs入口，传入的节点为查询节点target
	dfs(target, &ans, childrenDict)

	fmt.Println(ans)
	// 根据字典序进行排序
	sort.Strings(ans)

	// 逐行输出ans中的所有值，要注意ans中包含了查询target
	// 根据题意target是不应该输出的，需要多做一步判断
	for _, node := range ans {
		if node != target {
			fmt.Println(node)
		}
	}
}
