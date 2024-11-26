package dp

import (
    "fmt"
    "testing"
)

func TestJjsdj(t *testing.T) {
    //scanner := bufio.NewScanner(os.Stdin)
    //scanner.Scan()
    //str1 := scanner.Text()
    //scanner.Scan()
    //str2 := scanner.Text()
    //str1 := "abcdefghijklmnop"
    //str2 := "abcsafjklmnopqrstuvw"
    str1 := "nvlrzqcjltmrejybjeshffenvkeqtbsnlocoyaokdpuxutrsmcmoearsgttgyyucgzgcnurfbubgvbwpyslaeykqhaaveqxijc"
    str2 := "wkigrnngxehuiwxrextitnmjykimyhcbxildpnmrfgcnevjyvwzwuzrwvlomnlogbptornsybimbtnyhlmfecscmojrxekqmj"

    // 如果有多个结果时候要求输出较短的字符串里最先出现的结果，就去掉注释
    //// 保证str1是 短的字符串
    //if len(str1) > len(str2) {
    //    str1, str2 = str2, str1
    //}

    // dp[][]表示str1前i个字符串和str2前j个字符串组成的最大公共子串长度
    dp := make([][]int, len(str1)+1)
    for i := 0; i < len(str1)+1; i++ {
        dp[i] = make([]int, len(str2)+1)
    }
    maxLen, end := 0, 0

    for i := 1; i < len(str1)+1; i++ {
        for j := 1; j < len(str2)+1; j++ {
            if str1[i-1] == str2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = 0 // 不存在以i,j 结尾的最大公共子串
            }
            if dp[i][j] > maxLen {
                maxLen = dp[i][j]
                end = j
            }
        }
    }
    //fmt.Println(maxLen)
    fmt.Println(str2[end-maxLen : end])
}
