package od

import (
    "bufio"
    "fmt"
    "os"
    "testing"
)


func TestQuestion(t *testing.T)  {
    scanner := bufio.NewScanner(os.Stdin) // 标准输入输出
    scanner.Scan()
    str := scanner.Text()
    fmt.Println(str)
    wordLen := GetLastWordLen(str)
    fmt.Println(wordLen)

}

func GetLastWordLen(str string) (l int) {
    for i := len(str)-1; i > 0; i-- {
        if str[i] == ' '{
            return l
        }
        l++
    }
    return l
}