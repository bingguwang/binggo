// package main
//
// import (
//
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//
// )
//
// /*
// *
// 学校组织活动，将学生排成一个矩形方阵。
//
// 请在矩形方阵中找到最大的位置相连的男生数量。
//
// 这个相连位置在一个直线上，方向可以是水平的、垂直的、成对角线的或者反对角线的。
//
// 注：学生个数不会超过 10000。
//
// 输入描述
// 输入的第一行为矩阵的行数和列数，接下来的 n 行为矩阵元素，元素间用 , 分隔。
//
// 输出描述
// 输出一个整数，表示矩阵中最长的位置相连的男生个数。
// ————————————————
// 思路
//
// 典型的dfs
// */
// func main() {
//
//	// 输入
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	dims := strings.Split(scanner.Text(), ",")
//
//	rows, _ = strconv.Atoi(dims[0])
//	cols, _ = strconv.Atoi(dims[1])
//	matrix = make([][]string, rows)
//	for i := range matrix {
//		matrix[i] = make([]string, cols)
//	}
//	for i := 0; i < rows; i++ {
//		scanner.Scan()
//		elements := strings.Split(scanner.Text(), ",")
//		for j := 0; j < cols; j++ {
//			matrix[i][j] = elements[j]
//		}
//	}
//	// 初始化标记数组
//	visited = make([][]bool, cols*rows)
//	for i := 0; i < rows; i++ {
//		visited[i] = make([]bool, cols)
//	}
//	fmt.Println(matrix)
//	fmt.Println(visited)
//
//	var maxCount int
//	for i := 0; i < rows; i++ {
//		for j := 0; j < cols; j++ {
//			// 从左往右，从上到下
//			if matrix[i][j] == "M" && !visited[i][j] { // 是要找到对象，且未访问过
//				count := dfs(i, j)
//				if count > maxCount {
//					maxCount = count
//				}
//			}
//		}
//	}
//
//	fmt.Println("maxCount:", maxCount)
//
// }
//
// // 这里维护一个的记录是否访问过的次坐标点的标记数组
// var visited [][]bool
// var matrix [][]string
// var rows, cols int
//
// // 使用递归比较合适
// // 传入坐标，返回
//
//	func dfs(x, y int) int {
//		// 坐标超出边界，访问过了，或者不是要找的对象，则返回0
//		if x < 0 || x >= rows || y < 0 || y >= cols || visited[x][y] || matrix[x][y] != "M" {
//			return 0
//		}
//
//		visited[x][y] = true
//		count := 1
//		// 对传入的坐标，在四个方向上递归地搜索相连的'M'
//		directions := [][]int{{1, -1}, {0, 1}, {1, 0}, {1, 1}}
//		for _, dir := range directions {
//			nx, ny := x+dir[0], y+dir[1]
//			count += dfs(nx, ny)
//		}
//		return count
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取矩阵尺寸
	scanner.Scan()
	dims := strings.Split(scanner.Text(), ",")
	rows, _ := strconv.Atoi(dims[0])
	cols, _ := strconv.Atoi(dims[1])

	// 读取矩阵内容
	matrix := make([][]string, rows)
	for i := 0; i < rows; i++ {
		scanner.Scan()
		matrix[i] = strings.Split(scanner.Text(), ",") // 假设输入格式正确，只取第一个元素
	}
	fmt.Println(matrix)

	// 四个方向上的最长分别找出来
	// 但是有些没必要的计算，对于某一段连续的M，很显然我们考虑最边上M往另一边延申，考虑中间的M往两边延申，最后计算得到的连续的M的个数是一样的
	// 就比如 M M M F, 第一个M在横向得到的最大连续和第二个M得到的最大连续是一样的，而在计算第一个的最大连续的时候其实就遍历过M了，他们包含在同一段的连续里
	// 所以，这里就用一个标记数组
	check1 := make([][]bool, rows)
	// 同理其他方向也需要使用标记数组
	check2 := make([][]bool, rows)
	check3 := make([][]bool, rows)
	check4 := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		check1[i] = make([]bool, cols)
		check2[i] = make([]bool, cols)
		check3[i] = make([]bool, cols)
		check4[i] = make([]bool, cols)
	}

	var res = 0
	// 水平遍历
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 水平
			if matrix[i][j] == "M" && !check1[i][j] { // 加限制避免重复计算
				// (x,y)为横向移动时的坐标
				// (dx,dy)为横向移动的坐标偏差
				x := i
				y := j
				dy := 1
				// 当x和y不越界，且mat[x][y]仍为M时，进行横向移动的考虑
				// 修改check1[x][y]为已检查过，且(x,y)修改
				for x >= 0 && x < rows && y >= 0 && y < cols && matrix[x][y] == "M" {
					check1[x][y] = true
					y += dy // 向右移
				}
				// 退出while循环后，此段横向的连续的M的长度为y-j，更新答案
				if res < y-j {
					res = y - j
				}
			}

			//垂直
			if matrix[i][j] == "M" && !check2[i][j] { // 加限制避免重复计算
				// (x,y)为横向移动时的坐标
				// (dx,dy)为横向移动的坐标偏差
				x := i
				y := j
				dx := 1
				// 当x和y不越界，且mat[x][y]仍为M时，进行横向移动的考虑
				// 修改check1[x][y]为已检查过，且(x,y)修改
				for x >= 0 && x < rows && y >= 0 && y < cols && matrix[x][y] == "M" {
					check2[x][y] = true
					x += dx // 向下移
				}
				// 退出while循环后，此段横向的连续的M的长度为y-j，更新答案
				if res < x-i {
					res = x - i
				}
			}

			// 对角线
			if matrix[i][j] == "M" && !check3[i][j] { // 加限制避免重复计算
				// (x,y)为横向移动时的坐标
				// (dx,dy)为横向移动的坐标偏差
				x := i
				y := j
				dx, dy := 1, 1
				// 当x和y不越界，且mat[x][y]仍为M时，进行横向移动的考虑
				// 修改check3[x][y]为已检查过，且(x,y)修改
				for x >= 0 && x < rows && y >= 0 && y < cols && matrix[x][y] == "M" {
					check3[x][y] = true
					x += dx // 向下移
					y += dy // 向右移
				}
				// 退出while循环后，此段横向的连续的M的长度为 x - i，更新答案
				if res < x-i {
					res = x - i
				}
			}

			// 反对角线
			if matrix[i][j] == "M" && !check4[i][j] { // 加限制避免重复计算
				// (x,y)为横向移动时的坐标
				// (dx,dy)为横向移动的坐标偏差
				x := i
				y := j
				dx, dy := 1, -1
				// 当x和y不越界，且mat[x][y]仍为M时，进行横向移动的考虑
				// 修改check4[x][y]为已检查过，且(x,y)修改
				for x >= 0 && x < rows && y >= 0 && y < cols && matrix[x][y] == "M" {
					check4[x][y] = true
					x += dx // 向下移
					y += dy // 向左移
				}
				// 退出while循环后，此段横向的连续的M的长度为 x - i，更新答案
				if res < x-i {
					res = x - i
				}
			}
		}
	}

	// 输出结果
	fmt.Println(res)
}
