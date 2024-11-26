package leetCode

import (
    "fmt"
    "testing"
)

/**
  找出字符串数组里的最长公共前缀

初始化：最长前缀是第一个字符串

比较其他字符串 和 最长前缀 ，不断修改最长前缀，最后就是结果
*/
func TestGkas(t *testing.T) {
    strs := []string{"flower", "flow", "flight"}
    prefix := longestCommonPrefix(strs)
    commonPrefix := longestCommonPrefix([]string{"dog", "racecar", "car"})
    fmt.Println("res is :",prefix)
    fmt.Println("res is :",commonPrefix)
}
func longestCommonPrefix(strs []string) string {
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        prefix = getPrefix(prefix, strs[i])
    }
    return prefix
}

func getPrefix(a, b string) string {
    var s []byte
    i, j := 0, 0
    for i < len(a) && j < len(b) && a[i] == b[j] {
        s = append(s, a[i])
        i++
        j++
    }
    return string(s)
}
