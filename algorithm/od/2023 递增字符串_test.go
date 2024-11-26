package od

import (
    "fmt"
)

func minn(a, b int) int {
    if a < b {
        return a
    }
    return b
}

/**
dp[i][0] 表示在前 i 个字符中，以 A 结尾的严格递增子序列的最小修改次数
dp[i][1] 表示在前 i 个字符中，以 B 结尾的严格递增子序列的最小修改次数。

最终的答案是 dp[n][0] 和 dp[n][1] 中的最小值


*/

func minModification(s string) int {
    n := len(s)
    dp := make([][2]int, n+1)

    for i := 1; i <= n; i++ {
        if s[i-1] == 'A' { // 对于字符串中的第 i 个字符，如果它是 A
            // 可以把它添加到以 A 结尾的严格递增子序列中，不需要增加修改次数
            dp[i][0] = dp[i-1][0]
            // 或者把它修改成 B，并添加到以 B 结尾的严格递增子序列中
            dp[i][1] = minn(dp[i-1][0], dp[i-1][1]) + 1
        } else {
            // 可以把它添加到以 B 结尾的严格递增子序列中
            dp[i][0] = dp[i-1][0] + 1
            // 或者把它修改成 A，并添加到以 A 结尾的严格递增子序列中
            dp[i][1] = minn(dp[i-1][0], dp[i-1][1])
        }
    }

    return minn(dp[n][0], dp[n][1])
}

func main() {
    s := "ABABABBAA"
    fmt.Println(minModification(s))
}
