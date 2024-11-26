package doublepointer

import (
	"fmt"
	"testing"
)

/**
给你一个字符串 s，找到 s 中最长的回文子串。
*/
func TestSs(t *testing.T) {
	palindrome := longestPalindrome("abcbdd")
	palindrome2 := longestPalindrome2("abcbdd")
	fmt.Println(palindrome)
	fmt.Println(palindrome2)
}

func longestPalindrome(s string) string { // 中心扩散法
	res := ""
	for i := 0; i < len(s); i++ {
		res = getCenterSub(i, i, s, res)   // 中心是一个字符的时候
		res = getCenterSub(i, i+1, s, res) // 中心是两个字符的时候
	}
	return res
}

func getCenterSub(i, j int, s, res string) string {
	sub := ""

	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(sub) > len(res) {
		return sub
	}
	return res
}

func longestPalindrome2(s string) string { // DP
	res := ""
	dp := make([][]bool, len(s))
	// 初始化结果集
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ { // 以i为起点，j为终点的子串，看了个遍
			if s[i] != s[j] {
				dp[i][j] = false
			} else if j-i < 3 { // 小于等于3个字符
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
