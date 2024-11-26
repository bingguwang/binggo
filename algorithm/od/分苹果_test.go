package od

import (
	"fmt"
	"testing"
)

/**
把m个同样的苹果放在n个同样的盘子里，允许有的盘子空着不放，问共有多少种不同的分法？
注意：如果有7个苹果和3个盘子，（5，1，1）和（1，5，1）被视为是同一种分法。

*/

func partitionApples(m, n int) int {
	dp := make([][]int, m+1)
	//初始化
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][1] = 1
		dp[i][0] = 1 // 没有盘子，也算1
	}

	for j := 0; j < n+1; j++ {
		dp[1][j] = 1
		dp[0][j] = 1
	}
	// 已赋值的不要再去修改了
	for i := 2; i < m+1; i++ {
		for j := 2; j < n+1; j++ {
			// 共有两种情况 ,2种情况的并集就是总的可能
			// 如果至少有一个盘子空着，那么就相当于将i个苹果放入j-1个盘子中，即dp[i][j-1]；
			// 如果每个盘子都有至少一个苹果，那么就需要将i-j个苹果放入j个盘子中，即dp[i-j][j]。
			if i < j { // 苹果没有盘子多, 则属于至少有一个盘子空着
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i-j][j] + dp[i][j-1]
			}
		}
	}
	//for i:=0; i<m; i++ {
	//  for j:=0; j<=n;j++ {
	//      fmt.Println(dp[i][j])
	//  }
	//}

	return dp[m][n]
}
func TestUjsd(t *testing.T) {
	//apples := partitionApples(7, 3)
	apples := partitionApples(4, 1)
	fmt.Println(apples)
}
