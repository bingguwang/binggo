package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumSwap(1))
}
func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	arr := []rune(s)
	var stack []int
	for i := len(arr) - 1; i > 0; i-- {
		if len(stack) == 0 || arr[stack[len(stack)-1]] < arr[i] {
			stack = append(stack, i)
		}
	}
	fmt.Println("stack--", stack)
	for i := 0; i < len(arr); i++ {
		if len(stack) == 0 {
			break
		}
		if i < stack[len(stack)-1] {
			if arr[i] < arr[stack[len(stack)-1]] {
				arr[i], arr[stack[len(stack)-1]] = arr[stack[len(stack)-1]], arr[i]
				break
			}
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println(string(arr))
	atoi, _ := strconv.Atoi(string(arr))

	return atoi
}
