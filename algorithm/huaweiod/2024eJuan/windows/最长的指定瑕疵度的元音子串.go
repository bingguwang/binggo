package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	flaw, _ := strconv.Atoi(s.Text())
	s.Scan()
	str := s.Text()

	fmt.Println(handlecheckchar(str, flaw))
}

func handlecheckchar(str string, flaw int) int {
	n := len(str)

	// 收尾字符都是元音
	left := n
	// 找到第一个元音字母
	for i := 0; i < n; i++ {
		if checkchar(str[i]) {
			left = i
			break
		}
	}
	if left == n {
		return -1
	}
	var res int
	windowscountfuyin := 0
	// left已更新为第一个元音字母的位置
	for right := left; right < n; right++ {
		if checkchar(str[right]) { // 是元音
			if windowscountfuyin == flaw {
				// 更新最大长度
				res = max(right-left+1, res)
			}
		} else { // 辅音
			windowscountfuyin++
			// 是否需要移动左指针
			for left < n && (windowscountfuyin > flaw || !checkchar(str[left])) { // 因为最左边的要是元音
				if !checkchar(str[left]) {
					windowscountfuyin--
				}
				left++ // 左指针右移
			}
		}
	}
	fmt.Println(res)
	return res

}

// 是否是元音
func checkchar(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
