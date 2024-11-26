package od

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "testing"
    "unicode"
)

func TestJsd(t *testing.T) {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        chars := []rune(scanner.Text())
        // 其他字符
        otherChars := make([]bool, len(chars))
        // 字母集
        letters := []rune{}
        for i, c := range chars {
            // 判断是否是字母
            if unicode.IsLetter(c) {
                letters = append(letters, c)
                continue
            }
            otherChars[i] = true
        }
        sort.SliceStable(letters, func(i, j int) bool {
            // 转为小写再比较
            return unicode.ToLower(letters[i]) < unicode.ToLower(letters[j])
        })
        fmt.Println(string(letters))
        for i, c := range chars {
            if otherChars[i] { // 其他字符直接输出
                fmt.Printf("%c", c)
            } else {
                fmt.Printf("%c", letters[0]) // 字母一个一个拿出来去替换
                letters = letters[1:]
            }
        }
    }

}
