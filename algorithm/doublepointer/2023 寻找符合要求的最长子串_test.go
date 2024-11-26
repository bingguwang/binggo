package doublepointer

import (
    "fmt"
    "strings"
)

func main() {
    //text := "abbcacc"
    text := "ABACA123D"
    target := 'D'

    trim := strings.Split(text, string(target))
    for i := 0; i < len(trim); i++ {
        substring := lengthOfLongestSubstring1(trim[i])
        fmt.Println("res : ", substring)
    }

}

func lengthOfLongestSubstring1(s string) string {
    res := ""
    mp := map[byte]int{}
    left, right := 0, -1

    for left < len(s) { // 注意是左指针
        if right+1 < len(s) && mp[s[right+1]] <=  1 { // 已经有一个或0个时，还可以放
            mp[s[right+1]]++
            right++ // 右指针移动
        } else { // right的下一个元素是重复的
            mp[s[left]]-- // 因为左指针右移了，左指针原指的元素就要少计一个了
            left++        // 左指针右移
        }
        // 如果是要返回字符串
        if right-left+1 > len(res) {
            res = s[left : right+1]
            fmt.Println(s[left : right+1])
        }
    }
    return res
}