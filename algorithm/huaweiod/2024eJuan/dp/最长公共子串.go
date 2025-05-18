package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
func LCS(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}

	var maxsub int
	var endidx int // 只需要记录最长公共子串的末尾就行
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if maxsub < dp[i][j] {
				maxsub = dp[i][j]
				endidx = i
			}
		}
	}

	fmt.Println(str1[endidx-maxsub : endidx])

	return str1[endidx-maxsub : endidx]
}
