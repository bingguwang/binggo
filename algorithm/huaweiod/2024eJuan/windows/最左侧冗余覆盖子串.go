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
	s1 := s.Text()
	s.Scan()
	s2 := s.Text()
	s.Scan()
	k, _ := strconv.Atoi(s.Text())
	fmt.Print(handle(s1, s2, k))
}

func handle(s1, s2 string, k int) int {
	n1 := len(s1)
	n2 := len(s2)

	if n2 < n1+k {
		return -1 // 如果 s2 长度小于窗口长度，直接返回 -1
	}

	// 统计 s1 中每个字符的频率
	cnts1 := make(map[rune]int)
	for _, v := range s1 {
		cnts1[v]++
	}

	// 初始化窗口
	lenwin := n1 + k
	cnts2 := make(map[rune]int)
	for i := 0; i < lenwin; i++ {
		cnts2[rune(s2[i])]++
	}

	// 检查初始窗口是否满足条件
	if check(cnts1, cnts2) {
		return 0
	}

	// 滑动窗口
	for right := lenwin; right < n2; right++ {
		// 添加右边界字符
		cnts2[rune(s2[right])]++
		// 移除左边界字符
		left := right - lenwin
		cnts2[rune(s2[left])]--
		if cnts2[rune(s2[left])] == 0 {
			delete(cnts2, rune(s2[left])) // 一定记得删除
		}
		// 检查当前窗口是否满足条件
		if check(cnts1, cnts2) { // 找到第一个结果即可返回
			return left + 1
		}
	}

	// 如果没有找到满足条件的窗口，返回 -1
	return -1
}

// 检查 s2 的窗口是否覆盖 s1
func check(s1 map[rune]int, s2 map[rune]int) bool {
	for k, v := range s1 {
		if val, ok := s2[k]; !ok || val < v {
			return false
		}
	}
	return true
}
