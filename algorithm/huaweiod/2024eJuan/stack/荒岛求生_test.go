package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s2 := strings.Fields(s.Text())
	var nums []int
	for _, v := range s2 {
		t, _ := strconv.Atoi(v)
		nums = append(nums, t)
	}
	fmt.Println(nums)
	stack := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		for {
			if cur > 0 || len(stack) == 0 { // 向右的直接入栈,栈为空的直接入栈
				stack = append(stack, cur)
				break
			} else { // 向左的和栈顶比较大小
				topval := stack[len(stack)-1]
				if abs(topval) > abs(cur) { // 比栈顶小
					stack[len(stack)-1] = topval - abs(cur)
					break
				} else if abs(topval) == abs(cur) {
					stack = stack[:len(stack)-1]
					break
				} else { // 比栈顶大,更新血量
					cur = cur + topval // 更新当前要入栈的值
					stack = stack[:len(stack)-1]
					fmt.Println("更新后的cur:", cur)
				}
			}
		}
		fmt.Println("stack:", stack)
	}
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
