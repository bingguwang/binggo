package dp

import (
    "fmt"
    "testing"
)

/**
  每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
  明显的动态规划
1 1 1 1
f4=f3+1
f4=f2+1
类似斐波那契
*/
func TestVlim(t *testing.T) {
    fmt.Println("结果：", climbStairs(2))
    //fmt.Println("结果：", climbStairs(3))
    //fmt.Println("结果：", climbStairs(2))
}

func climbStairs(n int) int {
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-2] + dp[i-1]
    }
    return dp[n]
}


