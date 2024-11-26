package od

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestFS(t *testing.T) {
	var n int
	fmt.Scan(&n)
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		arr[i] = [2]int{}
		fmt.Scan(&arr[i][0], &arr[i][1])
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	var stack []int
	var count int
	for i := 0; i < len(text); i++ {
		if text[i] == '(' {
			continue
		} else if text[i] == ')' {
			// 出栈4个元素
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			y := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			_ = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			m := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			count += x * y * m
			stack = append(stack, m)
			stack = append(stack, x)
		} else { // 是字母，则把对应的矩阵的行， 列 入栈
			stack = append(stack, arr[text[i]-'A'][0])
			stack = append(stack, arr[text[i]-'A'][1])
		}
	}
	fmt.Println(count)
}

func pop(stack []int) int {
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return x
}
