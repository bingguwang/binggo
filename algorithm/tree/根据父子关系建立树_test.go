package tree

import (
	"fmt"
	"sort"
	"testing"
)

func TestName(t *testing.T) {
	var n int
	fmt.Scan(&n)

	// 使用二维切片来存储输入的无向图
	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		edges[i] = [2]int{a, b}
	}

	// 根据节点编号对边进行排序，以确保处理顺序的一致性
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][0] == edges[j][0] {
			return edges[i][1] < edges[j][1]
		}
		return edges[i][0] < edges[j][0]
	})

	// 初始化深度数组，节点编号从1开始，所以数组大小为n+1
	result := make([]int, n+1)
	for i := range result {
		result[i] = 1 // 初始化每个节点的深度为1（这里实际上是一个占位值，后续会更新）
	}

	// 遍历边来更新深度
	for _, edge := range edges {
		parent, child := edge[0], edge[1]
		result[child] = result[parent] + 1
	}

	// 找出最大深度
	maxHeight := 0
	for _, depth := range result[1:] { // 忽略节点0，因为题目从1开始
		if depth > maxHeight {
			maxHeight = depth
		}
	}

	fmt.Println(maxHeight)
}
