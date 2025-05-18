package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var (
	maxHeight  int
	minSteps   int
	directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
)

func dfsaa(matrix [][]int, visited [][]bool, point Point, steps int, k int) {
	height := matrix[point.x][point.y]

	// 更新最高高度和最短步数
	if height > maxHeight || (height == maxHeight && steps < minSteps) {
		maxHeight = height
		minSteps = steps
	}

	// 标记当前位置为已访问
	visited[point.x][point.y] = true

	// 检查四个相邻位置
	for _, dir := range directions {
		newX := point.x + dir.x
		newY := point.y + dir.y
		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) && !visited[newX][newY] {
			heightDiff := matrix[newX][newY] - height
			if heightDiff == k || heightDiff == -k {
				dfsaa(matrix, visited, Point{x: newX, y: newY}, steps+1, k)
			}
		}
	}

	// 回溯，取消标记当前位置为已访问
	visited[point.x][point.y] = false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取矩阵的行数 m 和列数 n
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, " ")
	m, _ := strconv.Atoi(parts[0])
	n, _ := strconv.Atoi(parts[1])

	// 读取矩阵中的高度值
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		row := scanner.Text()
		heights := strings.Split(row, " ")
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j], _ = strconv.Atoi(heights[j])
		}
	}

	// 读取允许的高度差 k
	scanner.Scan()
	kStr := scanner.Text()
	k, _ := strconv.Atoi(kStr)

	// 初始化访问数组
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 初始化最高高度和最短步数
	maxHeight = 0
	minSteps = 0

	// 从起点 (0, 0) 开始 DFS
	dfsaa(matrix, visited, Point{x: 0, y: 0}, 0, k)

	// 输出结果
	fmt.Println(maxHeight)
	fmt.Println(minSteps)
}
