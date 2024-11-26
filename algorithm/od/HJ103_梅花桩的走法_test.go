package od

import (
    "fmt"
    "testing"
)

func TestGsd(t *testing.T) {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    // 初始化，每个桩子最少可以走1步
    dp := make([]int, n)
    for i := 0; i < len(dp); i++ {
        dp[i] = 1
    }
    for i := 0; i < len(arr); i++ {
        for j := 0; j < i; j++ { // 这样迭代，以满足所有的结果都能遍历到
            if arr[j] < arr[i] { // i是终点，因为只能从低到高，所以j要比i小
                dp[i] = max(dp[i], dp[j]+1) // 状态转移方程
            }
        }
    }
    // 取最大值
    maxVal := 0
    for i := 0; i < len(dp); i++ {
        if maxVal < dp[i] {
            maxVal = dp[i]
        }
    }
    fmt.Println(maxVal)
}