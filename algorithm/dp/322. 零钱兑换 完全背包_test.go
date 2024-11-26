package dp

import (
"fmt"
)

func main() {
    coins := []int{1, 2, 5}
    amount := 11

    // dp[i]表示凑出i所需要的最少硬币数
    dp := make([]int, amount+1)
    for i := 0; i < amount+1; i++ {
        dp[i] = amount + 1
    }
    // 初始化
    dp[0] = 0
    for i := 0; i < amount+1; i++ {
        for j := 0; j < len(coins); j++ {
            if coins[j] <= i { // 可选可不选
                dp[i] = min(dp[i], dp[i-coins[j]]+1)
            }else {
                dp[i] = dp[i] // 不能选这个硬币， 可见是可以省去的
            }
        }
    }

    fmt.Println(dp)
    fmt.Println(dp[amount])
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}

