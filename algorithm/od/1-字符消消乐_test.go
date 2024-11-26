package od

import (
    "fmt"
    "testing"
)

/**
  使用栈实现
*/
func TestAksaad(t *testing.T) {
    fmt.Println(XiaoXiaoString("abbaca"))
    fmt.Println(XiaoXiaoString("abbaaccb"))
}

/*
abbbaca
ab
b
*/
func XiaoXiaoString(str string) string {
    var stack []byte
    var top byte
    fmt.Println(str,"消除过程是：")
    for i := 0; i < len(str); i++ {
        if len(stack) == 0 {
            stack = append(stack, str[i])
            top = stack[0]
            continue
        }
        if top == str[i] {
            continue
        } else if stack[len(stack)-1] == str[i] {
            top = stack[len(stack)-1]
            // 出栈
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, str[i])
        }
        fmt.Println(string(stack))
    }
    return string(stack)
}
