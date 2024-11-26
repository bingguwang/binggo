package doublepointer

import (
    "fmt"
    "testing"
)

/**
  在一个字符串里寻找没有重复字母的最长子串。
*/

func TestSer(t *testing.T) {
    substring := lengthOfLongestSubstring("abccccccc")
    fmt.Println(substring)
}

// 双指针
func lengthOfLongestSubstring(s string) int {
    res := 0
    mp := map[byte]int{}
    left, right := 0, -1

    for left < len(s) { // 注意是左指针
        if right+1 < len(s) && mp[s[right+1]] == 0 { // right的下一个元素没重复
            mp[s[right+1]]++
            right++ //右指针移动
        } else { // right的下一个元素是重复的
            mp[s[left]]-- // 因为左指针右移了，左指针原指的元素就要少计一个了
            left++        // 左指针右移
        }
        res = max(res, right-left+1)
        /** 如果是要返回字符串
          if right-left+1 > len(res) {
               res = str[left : right+1]
           }
        */
    }
    return res
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
