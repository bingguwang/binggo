package od

import (
    "fmt"
    "testing"
)

func TestHsawwq(t *testing.T) {

    /**
          对“dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作。”的补充理解：

      以 word1 为 "horse"，word2 为 "ros"，且 dp[5][3] 为例，即要将 word1的前 5 个字符转换为 word2的前 3 个字符，也就是将 horse 转换为 ros，因此有：

      (1) dp[i-1][j-1]+1，即先将 word1 的前 4 个字符 hors 转换为 word2 的前 2 个字符 ro，然后将第五个字符 word1[4]（因为下标基数以 0 开始） 由 e 替换为 s（即替换为 word2 的第三个字符，word2[2]）

      (2) dp[i][j-1]+1，即先将 word1 的前 5 个字符 horse 转换为 word2 的前 2 个字符 ro，然后在末尾补充一个 s，即插入操作

      (3) dp[i-1][j]+1，即先将 word1 的前 4 个字符 hors 转换为 word2 的前 3 个字符 ros，然后删除 word1 的第 5 个字符
    */

    var a, b string
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Println(a)
    fmt.Println(b)

    var dp [][]int
    dp = make([][]int, len(a)+1)
    for i := 0; i < len(a)+1; i++ {
        dp[i] = make([]int, len(b)+1)
    }
    // 初始化
    for i := 1; i < len(a)+1; i++ {
        dp[i][0] = dp[i-1][0] + 1
    }
    for i := 1; i < len(b)+1; i++ {
        dp[0][i] = dp[0][i-1] + 1
    }
    // dp[i][j]表示0-i 变为b的0-j的距离
    /**
      变换只有三种，a新增1个子串，a替换一个子串, a删除一个子串
    */
    fmt.Println(dp)
    for i := 1; i < len(a)+1; i++ {
        for j := 1; j < len(b)+1; j++ {
            if a[i-1] == b[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min([]int{dp[i-1][j-1], dp[i-1][j], dp[i][j-1]}) + 1
            }
        }
    }
    fmt.Println(dp[len(a)][len(b)])
}
func min(a []int) int {
    var mi int
    for i := 0; i < len(a); i++ {
        if mi > a[i] {
            mi = a[i]
        }
    }
    return mi
}
