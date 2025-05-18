package od

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

/**
给定一个可存储若干单词的字典，找出指定单词的所有相似单词，并且按照单词名称从小到大排序输出。
单词仅包括字母，但可能大小写并存（大写不一定只出现在首字母）。
相似单词说明：给定一个单词X，如果通过任意交换单词中字母的位置得到不同的单词Y，那么定义Y是X的相似单词，
如abc、bca即为相似单词（大小写是不同的字母，如a和A算两个不同字母）。
字典序排序： 大写字母<小写字母。同样大小写的字母，遵循26字母
6
abc
dasd
tadasd
bcdasd
bcda
cbda
Abcd
*/
func main() {
    var n int
    fmt.Scan(&n)

    var str []string
    scanner := bufio.NewScanner(os.Stdin)
    for i := 0; i < n; i++ {
        scanner.Scan()
        text := scanner.Text()
        str = append(str, text)
    }
    scanner.Scan()
    word := scanner.Text()
    var res []string
    for i := 0; i < len(str); i++ {
        if checkIsSimilar(word, str[i]) { // 是相似单词
            res = append(res, str[i])
        }
    }
    sort.Strings(res)
    if len(res) == 0 {
        fmt.Println("null")
    } else {
        fmt.Println(res)
    }
}
func checkIsSimilar(a, b string) bool {
    strs1 := strings.Split(a, "")
    sort.Strings(strs1)
    a = strings.Join(strs1, "")
    strs2 := strings.Split(b, "")
    sort.Strings(strs2)
    b = strings.Join(strs2, "")

    return a == b
}
