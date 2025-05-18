package main

import "fmt"

/**
数组的遍历


维护一个 visited 数组来记录已访问的节点，以避免重复访问和陷入无限循环

*/

type Graph struct {
	Vertices int
	Edges    map[int][]int // 邻接表，邻接点和边的关系，其中key 是邻接表的表头，一般是邻接点的编号, value是个数组，表示和这个连接点相邻接的所有节点
}

func (graph *Graph) dfsGraph(start int, visited []bool) {
	// 标记当前节点为已访问
	visited[start] = true

	// 访问当前节点
	processNode(start)

	// 获取当前节点的所有邻接节点
	for _, neighbor := range graph.Edges[start] {
		if !visited[neighbor] { // 没访问过的邻接点继续递归访问
			graph.dfsGraph(neighbor, visited)
		}
	}
}

func processNode(node int) {
	// 在这里处理节点，比如打印节点的值
	fmt.Println(node)
}

func (graph *Graph) dfsUtils(start int, visited map[int]bool) {
	// 标记当前节点为已访问
	visited[start] = true

	// 访问当前节点
	processNode(start)

	// 获取当前节点的所有邻接节点
	for _, neighbor := range graph.Edges[start] {
		if !visited[neighbor] { // 没访问过的邻接点继续递归访问
			graph.dfsUtils(neighbor, visited)
		}
	}
}

// 计算连通分量个数
func (g *Graph) CountConnectedComponents() int {
	visited := make(map[int]bool)
	count := 0

	for node := range g.Edges {
		if !visited[node] { // 没有访问的节点可以拿来递归
			g.dfsUtils(node, visited)
			count++
		}
	}

	return count
}

// 使用示例
func main() {
	graph := &Graph{
		Vertices: 5,
		Edges: map[int][]int{
			0: {1, 2},
			1: {0, 3, 4},
			2: {0},
			3: {1},
			4: {1},
		},
	}

	// 初始化标记节点
	visited := make([]bool, graph.Vertices)
	for i := range visited {
		visited[i] = false
	}

	// 从节点 0 开始进行 DFS
	graph.dfsGraph(0, visited)

	fmt.Println("连通分量---", graph.CountConnectedComponents())

}
