package dp

import "fmt"

/**
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

dp[i]表示 leetcode 的0到i-1个字符串可以用 wordDict 里的单词表示
那我们可以枚举 j = 0, 1, ..., i-1，并检查 dp[j] 是否为真，以及 s[j+1, i]（即从位置 j+1 到位置 i 的子串）是否出现在字典中。如果两个条件都满足，那么 dp[i] 为真
dp[0] = true
*/
func main() {
    s := "leetcode"
    wordDict := []string{"leet", "code"}
    dict := wordBreak(s, wordDict)
    fmt.Println(dict)
}
func wordBreak(s string, wordDict []string) bool {
    mp := make(map[string]bool)
    for _, v := range wordDict {
        mp[v] = true
    }
    dp := make([]bool, len(s)+1)
    dp[0] = true

    for i := 1; i < len(s)+1; i++ {
        for j := 0; j < i; j++ { // j就是 0到i-1这段字符串里的分割点
            if dp[j] && mp[s[j:i]] { // 0到j-1  j到i-1
                dp[i] = true // 找到一个分割点，可以使得0到i-1这段字符可以被分割为两个在结合里的单词，就不用找了，break
            }
        }
    }
    fmt.Println(dp)
    return dp[len(s)]
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
