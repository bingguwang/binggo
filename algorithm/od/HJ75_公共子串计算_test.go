package od

import "fmt"

// 注：子串的定义指一个字符串删掉其部分前缀和后缀（也可以不删）后形成的字符串。
// 也就是子串必须是连续字符组成，所以使用DP时和LeetCode那边的解法有不同

func getSub2(a, b string) {
    dp := make([][]int, len(a)+1)
    for i := 0; i < len(a)+1; i++ {
        dp[i] = make([]int, len(b)+1)
    }

    for i := 0; i < len(a)+1; i++ {
        dp[i][0] = 0
    }
    for i := 0; i < len(b)+1; i++ {
        dp[0][i] = 0
    }
    var maxLen int
    for i := 1; i <= len(a); i++ {
        for j := 1; j <= len(b); j++ {
            if a[i-1] == b[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = 0 // 至此最大值已经被存在maxLen里了，后面的需要重新计算, 因为这里和LeetCode里不太一样，OD这里的要求是连续字符组成的子串
            }
            if maxLen < dp[i][j] {
                maxLen = dp[i][j]
            }
        }

    }
    fmt.Println(maxLen)
}