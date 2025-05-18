package main

import (
	"fmt"
	"math"
)

func main() {
	//handle("ABCDEFG", 3)
	handle("ABCDEFGHIJKLMN", 4)
}

func handle(str string, n int) {
	length := len(str)
	m := int(math.Ceil(float64(length) / float64(n)))
	var arr = make([][]byte, n)
	for i := range arr {
		arr[i] = make([]byte, m)
	}

	i, j := 0, 0
	isdown := true
	for _, v := range str {
		arr[i][j] = byte(v)
		// fmt.Println("添加", string(arr[i][j]))

		if isdown {
			newi := i + 1
			if newi == n { // 超过下边界
				j++
				isdown = false
			} else {
				i = newi
			}
		} else {
			newi := i - 1
			if newi < 0 { // 超过上边界
				j++
				isdown = true
			} else {
				i = newi
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] != 0 {
				fmt.Printf("%c", arr[i][j])
			}
		}
		fmt.Println()
	}

}
