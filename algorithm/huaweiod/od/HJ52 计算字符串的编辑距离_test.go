package od

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

//`dp[i][j]` 代表 `word1` 到 `i` 位置转换成 `word2` 到 `j` 位置需要最少步数
func TestName(t *testing.T) {
	//minDistance("intention", "execution")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text1 := scanner.Text()
	scanner.Scan()
	text2 := scanner.Text()
	minDistance(text1, text2)
}

func minDistance(word1, word2 string) {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	// 第一行
	for j := 1; j < n+1; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}
	// 第一列
	for i := 1; i < m+1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] { // 当前字符相同
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + getMin(getMin(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[m][n])
	for _, v := range dp {
		fmt.Println(v)
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
