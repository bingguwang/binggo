package od

import (
	"fmt"
	"testing"
)

func TestKlsd(t *testing.T) {
	Walk(2, 2)
}

func Walk(n, m int) {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	// 初始化
	dp[1][1] = 2
	for i := 0; i < n+1; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < m+1; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}

}
