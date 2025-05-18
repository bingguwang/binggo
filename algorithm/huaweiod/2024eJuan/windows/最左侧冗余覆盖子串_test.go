package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
*
给定两个字符串Q s1 和 s2 和正整数 k，其中 s1 长度为 n1 ，s2 长度为 n2 。在 s2 中选一个子串，若满足下面条件，则称 s2 以长度 k 冗余覆盖 s1该子串长度为 n1 +k
该子串中包含 s1 中全部字母
该子串每个字母出现次数不小于 s1 中对应的字母给定 s1，s2，k，求最左侧的 s2 以长度 k 冗余覆盖 s1的子串的首个元素的下标，如果没有返回 -1。
举例：
s1 = "ab"
s2 = "aabcd"
k = 1
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s1 := scanner.Text()
	scanner.Scan()
	s2 := scanner.Text()
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	// 子串长度
	lenwin := k + len(s1)

	for i := 0; i < len(s2) && i+lenwin < len(s2); i++ {
		if checksub(s1, s2[i:i+lenwin]) {
			fmt.Println(i)
			fmt.Println(s2[i : i+lenwin])
			return
		}
	}
	fmt.Println(-1)
}

func checksub(s1, substr string) bool {
	mp := make(map[rune]int)
	for _, v := range substr {
		mp[v]++
	}
	for _, v := range s1 {
		if val, ok := mp[v]; !ok || val == 0 {
			return false
		} else {
			mp[v]--
		}
	}
	return true
}
