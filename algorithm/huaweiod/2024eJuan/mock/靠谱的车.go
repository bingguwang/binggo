package main

import (
	"fmt"
	"strconv"
)

/*
*
司机的计费表相当于一个“无 4 十进制”系统：
在正常十进制中，每一位的可能值是 0~9。
在司机的计费表中，每一位的可能值是 0, 1, 2, 3, 5, 6, 7, 8, 9（去掉了 4）。
这种规则类似于一个 9 进制系统，但它的基数并不是完全均匀分布的，而是跳过了 4。
从司机计费表到实际费用的转换：
将司机的计费表读数视为一个“无 4 十进制”数。
将其转换为正常的十进制数即可得到实际费用。
*/

// Scanner 函数
func Scanner(N int) int {
	NStr := strconv.Itoa(N) // 将整数 N 转换为字符串
	sum := 0
	fmt.Println(NStr)
	for i := len(NStr) - 1; i >= 0; i-- {
		// 获取字符串从右往左的第 i 个字符，并转换为整数
		NInt, _ := strconv.Atoi(string(NStr[i]))
		// 如果数字大于 4，则减 1
		// 0-3 5-9 九个数，所以是9进制，大于4的时候数值先减掉1就是真正的9进制了
		if NInt > 4 {
			NInt -= 1
		}
		// 计算当前位的值并累加到 sum
		fmt.Println("指数--", len(NStr)-1-i)
		sum += NInt * pow(9, len(NStr)-1-i)
	}
	fmt.Println(sum)
	return sum
}

// 辅助函数：计算幂次方
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func main() {
	// 测试 Scanner 函数
	//Scanner(5)   // 示例输入 1
	//Scanner(17)  // 示例输入 2
	Scanner(100) // 示例输入 3
}
