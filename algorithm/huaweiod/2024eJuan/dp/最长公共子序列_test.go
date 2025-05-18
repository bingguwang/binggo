package main

import "fmt"

func main() {
	//lcs := longestCommonSubsequence("1A2C3D4B56", "B1D23A456A")
	lcs := longestCommonSubsequence("abcde", "ace")
	fmt.Println(lcs)
}

// 难度升级， 需要求子串而不仅仅是长度
func LCS(s1 string, s2 string) string {
	return ""
}

func longestCommonSubsequence(text1 string, text2 string) int {
	/**
	dp[i][j] 表示字符串text1的前i个字符串，也就是 0-i-1这些字符, text2的前j个字符串，也就是 0- j-1这些字符里，最长公共子序列长度
	if text1[i] == text2[j]
	dp[i][j] = dp[i-1][j-1]+1
	else
	dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	*/

	var dp = make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	//dp[0][0] = 0
	//dp[1][0] = 0
	//dp[0][1] = 0
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}

	// 如果要升级难度，不仅要知道长度而已，而且要知道结果的子串
	// 那就反推
	// 反推状态转移公式得到最长公共子序列
	var res string
	for i, j := len(text1), len(text2); dp[i][j] > 0; {
		if text1[i-1] == text2[j-1] { //反推公式中相等的场景
			// 该值一定是被选取到的，根据状态转移公式，当时两条字符串的下标都前进一位
			res = string(text1[i-1]) + res // 合并时注意添加进res前面
			i, j = i-1, j-1                // 反推公式，两个下标都后退一位
		} else if dp[i-1][j] >= dp[i][j-1] { // 当时dp[i][j]结果来自dp[i-1][j]
			i = i - 1 // 反推公式，行下标即text1下标i后退一位
		} else if dp[i-1][j] < dp[i][j-1] { // 当时dp[i][j]结果来自dp[i][j-1]
			j = j - 1 // 反推公式，列下标即text2下标j后退一位
		}
	}
	fmt.Println(res)

	return dp[len(text1)][len(text2)]
}
