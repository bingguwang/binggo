package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, -2},
	}
	endpointx := 10
	arr = append(arr, []int{endpointx, 0}) // 因为球面积是直接用cury就行了，所以知道前面的宽就行了
	fmt.Println(draw(arr))
}

func draw(code [][]int) int {
	var res int
	curx, cury := 0, 0
	for i := 0; i < len(code); i++ {
		fmt.Println("当前坐标:", curx, " ", cury)
		nxtx, nxty := code[i][0], cury+code[i][1]
		fmt.Println("下一个坐标:", nxtx, " ", nxty)

		// 面积计算
		area := abs((nxtx - curx) * cury)
		fmt.Println("面积是:", area)
		res += area
		curx, cury = nxtx, nxty
	}
	fmt.Println(curx, " ", cury)

	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
