package main

import (
	"fmt"
	"math"
)

func main() {
	// 定义 n 和 m，可以修改为从控制台读取输入
	n := 9
	m := 4
	columnCount := int(math.Ceil(float64(n) / float64(m))) // 计算列数

	// 初始化矩阵
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, columnCount)
	}

	// 填充螺旋矩阵
	fillSpiral(matrix, n)

	// 输出矩阵
	for i := 0; i < m; i++ {
		for j := 0; j < columnCount; j++ {
			if matrix[i][j] == 0 {
				fmt.Printf("%2s ", "*") // 如果值为 0，用 "*" 表示
			} else {
				fmt.Printf("%2d ", matrix[i][j]) // 否则输出数字
			}
		}
		fmt.Println()
	}
}

// fillSpiral 按照螺旋顺序填充矩阵
func fillSpiral(matrix [][]int, total int) {
	// 将要填充的数字，从1开始
	num := 1

	// 方向变量，四个元素分别表示：向右，向下，向左，向上
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// 初始化方向变量的索引，0表示初始方向为向右
	dir := 0

	// 初始化行列索引
	x, y := 0, 0

	// 继续填充，直到所有的数字被填充完
	for num <= total {
		// 将数字填充到矩阵的当前位置，然后数字加1
		matrix[x][y] = num
		num++

		// 尝试按照当前方向去获取新的行列索引
		newX := x + directions[dir][0]
		newY := y + directions[dir][1]

		// 检查新的行列索引是否超出矩阵的边界，或者该位置已经被填充过
		// 如果是，则将方向变量的索引加1（取模为了形成循环：向右，向下，向左，向上）
		if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[0]) || matrix[newX][newY] != 0 {
			dir = (dir + 1) % 4
			// 类似贪吃蛇，向右不行就向下，向下不行就向左，向左不行就向上，向上不行就向右
		}
		// 确定好了往哪个方向进行下一步之后
		// 根据新的方向更新行列索引
		x += directions[dir][0]
		y += directions[dir][1]
	}
}
