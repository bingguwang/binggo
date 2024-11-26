package dp

import (
    "fmt"
    "strings"
    "testing"
)

// 不需要子串是连续字符组成的，只要是
//一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

func TestJaqd(t *testing.T) {
    //sub := getMaxSub("abcdefgjjjjjjjjjhijklmnop", "abcsafgjjjjjjjjjjjjklmnopqrstuvw")
    //fmt.Println(sub)
    //getSub3("abcdefgjjjjjjjjjhijklmnop", "abcsafgjjjjjjjjjjjjklmnopqrstuvw")
    getSub3("abcdefghijklmnop", "abcsafjklmnopqrstuvw")
}

/**
  abcdefghijklmnop
  abcsafjklmnopqrstuvw

    暴力法
*/
func getMaxSub(a, b string) string {
    var longer string
    var shorter string
    if len(shorter) > len(longer) {
        longer, shorter = shorter, longer
    }
    maxSub := ""
    for i := 0; i < len(shorter); i++ { // 子串的其实字符串的位置
        for k := i + 1; k <= len(shorter); k++ { // 子串长度
            sub := shorter[i:k] // 子串
            //for j := 0; j < len(longer); j++ {
            if strings.Index(longer, sub) >= 0 {
                if len(maxSub) < len(sub) {
                    maxSub = sub
                }
                break
            }
            //}
        }
    }
    return maxSub
}

/**
  DP
  dp[i][j] 表示2个字符串分别在取0-i,0-j部分时最大的公共子序列的长度
*/

func getSub3(a, b string) {
    dp := make([][]int, len(a)+1)
    for i := 0; i < len(a)+1; i++ {
        dp[i] = make([]int, len(b)+1)
    }
    // dp[i][j]序列a的前i个元素和序列b的前j个元素的最长公共子序列的长度。
    // dp[i][j] = dp[i-1][j] + 1
    var maxl int
    for i := 1; i < len(a)+1; i++ {
        for j := 1; j < len(b)+1; j++ {
            if a[i-1] == b[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max22(dp[i-1][j], dp[i][j-1])
            }
            if maxl < dp[i][j] { // 最大值保存
                maxl = dp[i][j]
            }
        }
    }
    for _, v := range dp {
        fmt.Println(v)
    }
    fmt.Println(maxl)
}
func max22(i, j int) int {
    if i < j {
        return i
    }
    return j
}
