package first

import (
    "fmt"
    "testing"
)
/**
给出一个字符串 s（仅含有小写英文字母和括号）。
请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
注意，您的结果中 不应 包含任何括号。

输入：s = "(ed(et(oc))el)"
输出："leetcode"
解释：先反转子字符串 "oc" ，接着反转 "etco" ，然后反转整个字符串。

思路 一眼就知道要用栈
 */
func TestFJBSAQ(t *testing.T) {
    duplicate := reverseParentheses2("(ed(et(oc))el)")
    fmt.Println(duplicate)
}

func reverseParentheses2(s string) string {
    stk := []string{}
    str := ""
    for _, ch := range s {
        if ch == '(' {
            stk = append(stk, str) // 单词入栈
            str = ""
        } else if ch == ')' {
            str = reverseStr([]byte(str)) // 先翻转
            // 栈顶出栈，并接在str左边
            str = stk[len(stk)-1] + str
            stk = stk[:len(stk)-1]
        } else {
            str += string(ch) // 接在str右边
        }
    }
    return str
}
func reverseStr(str []byte) string {
    i, j := 0, len(str)-1
    for i < j {
        str[i], str[j] = str[j], str[i]
        j--
        i++
    }
    return string(str)
}
