package main

import (
	"fmt"
)

// 定义栈元素类型
type StackElement struct {
	count int // 重复次数
	pos   int // 当前字符串的起始位置
}

func decodeString(s string) string {
	var ans string
	stk := []StackElement{} // 使用切片模拟栈
	count := 0

	for _, x := range s {
		if isDigit(x) {
			// 处理数字
			count = 10*count + int(x-'0')
		} else if x == '[' {
			// 将当前计数和字符串长度压入栈
			stk = append(stk, StackElement{count, len(ans)})
			count = 0
		} else if isAlpha(x) {
			// 处理字母，直接加入结果字符串
			ans += string(x)
		} else if x == ']' {
			// 弹出栈顶元素并进行字符串重复操作
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1] // 弹出栈顶

			n := top.count
			str := ans[top.pos:] // 获取需要重复的部分
			// 注意因为ans里已经加过一次sub，所以这里应该是再拼接repeat-1次即可
			for i := 0; i < n-1; i++ {
				ans += str
			}
		}
	}

	return ans
}

// 判断字符是否为数字
func isDigit(x rune) bool {
	return x >= '0' && x <= '9'
}

// 判断字符是否为字母
func isAlpha(x rune) bool {
	return (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')
}

func main() {
	// 测试示例
	input := "2[a2[bc]]"
	output := decodeString(input)
	fmt.Println(output) //
}
