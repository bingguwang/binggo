package main

import "fmt"

func longestPalindrome(s string) string {
	lenS := len(s)
	if lenS < 2 { // 长度小于2
		return s
	}

	dp := make([][]bool, lenS)
	for i := range dp {
		dp[i] = make([]bool, lenS)
	}

	result := s[0:1] // 至少包含一个字符
	for i := 0; i < lenS; i++ {
		dp[i][i] = true // 单个字符都是回文
	}

	for length := 2; length <= lenS; length++ { // 子串长度从2开始逐步增加, 最大的子串长度就是字符串的总长度
		for start := 0; start < lenS-length+1; start++ { // 起始下标一个个遍历
			end := start + length - 1 // 结尾下标
			if s[start] != s[end] {
				dp[start][end] = false
			} else {
				if length == 2 || dp[start+1][end-1] { // 长度是2或者子串是回文
					dp[start][end] = true
					if length > len(result) {
						result = s[start : end+1]
					}
				}
			}
		}
	}
	fmt.Println(dp)
	return result
}

func main() {
	fmt.Println(longestPalindrome("babad")) // 示例调用
}
