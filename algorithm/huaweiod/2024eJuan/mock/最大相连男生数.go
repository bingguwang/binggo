package main

import (
	"fmt"
)

var DIRECTIONS = [4][2]int{
	{0, 1},  // dx =0  dy =1 横向
	{1, 0},  // dx =1  dy =0 竖向
	{1, 1},  // dx =1  dy =1 对角线
	{1, -1}, // dx =1  dy =-1 饭对角线
}

func maxConsecutiveBoys(mat [][]string, n, m int) int {
	// 初始化答案变量，指的是最长连续男生数量
	ans := 0

	// 四个检查矩阵，check[0]、check[1]、check[2]、check[3]
	// 分别表示横向、纵向、对角线、反对角线方向上
	// 某一个M是否已经在某段连续的M中被考虑过
	check := make([][][]bool, 4)
	for k := 0; k < 4; k++ {
		check[k] = make([][]bool, n)
		for i := 0; i < n; i++ {
			check[k][i] = make([]bool, m)
		}
	}

	// 从上到下，从左到右顺序枚举矩阵中的每一个元素
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 考虑横向、纵向、主对角线方向、反对角线方向4个方向
			for k := 0; k < 4; k++ {
				// 如果某个M尚未被考虑包含在某段连续的的M中
				// 则考虑这段M
				if mat[i][j] == "M" && check[k][i][j] == false {
					// (x,y)为移动时的坐标
					// (dx,dy)为移动的坐标偏差
					x, y := i, j
					dx, dy := DIRECTIONS[k][0], DIRECTIONS[k][1]

					// 当x和y不越界，且mat[x][y]仍为M时，进行进一步移动的考虑
					// 修改check[k][x][y]为已检查过，且(x,y)修改
					for x >= 0 && x < n && y >= 0 && y < m && mat[x][y] == "M" {
						check[k][x][y] = true
						x += dx
						y += dy
					}

					// 退出while循环后，此段连续的M的长度为x-i或y-j（取其中较大值），更新答案
					ans = max(ans, max(x-i, y-j))
				}
			}
		}
	}

	fmt.Println(ans)
	return ans
}
