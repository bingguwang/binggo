package dp

import (
    "fmt"
    "testing"
)

func TestJkam(t *testing.T) {
    //stairs := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
    stairs := minCostClimbingStairs([]int{10, 15, 20})
    fmt.Println(stairs)
}
func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost)+1)
    dp[0] = 0
    dp[1] = 0
    // 从0开始
    for i := 2; i < len(cost)+1; i++ {
        dp[i] = Min(cost[i-1]+dp[i-1], cost[i-2]+dp[i-2])
        fmt.Println(dp)
    }
    return dp[len(cost)]
}
func Min(i, j int) int {
    if i < j {
        return i
    }
    return j
}
