package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// 创建读取器以读取输入
	reader := bufio.NewReader(os.Stdin)

	// 读取字符 a 和字符串 s
	a, _ := reader.ReadString('\n')
	a = string(a[0]) // 只取第一个字符
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	// 初始化变量
	var ans int
	hashTable := make([]int, 128) // 数组记录窗口中字符的频数
	left := 0
	// 遍历字符串 s
	for right, ch := range s {
		if ch != rune(a[0]) { // 如果当前字符不为应该排除的字符，则进行滑窗
			hashTable[ch]++ // A1: 更新当前字符的频数

			// A2: 如果某个字符的频数达到 3，则缩小窗口
			for hashTable[ch] == 3 {
				leftChar := s[left]
				hashTable[leftChar]--
				left++
			}

			// A3: 更新最大子串长度
			ans = max(ans, right-left+1)
		} else { // 如果当前字符为应该排除的字符，则清空滑窗
			left = right + 1
			hashTable = make([]int, 128) // 重置哈希表
		}
	}

	// 输出结果
	fmt.Println(ans)
}

// 辅助函数：返回两个整数的最大值
func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
