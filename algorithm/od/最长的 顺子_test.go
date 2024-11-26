package od

import (
    "fmt"
    "strings"
    "testing"
)

func TestGjaslw(t *testing.T) {
    var str1, str2 string
    fmt.Scan(&str1)
    fmt.Scan(&str2)
    var remain [18]int
    for i := 3; i <= 14; i++ {
        remain[i] = 4
    }
    fmt.Println(remain)
    strs1 := strings.Split(str1, "-")
    strs2 := strings.Split(str2, "-")
    mp := map[string]int{"3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13, "A": 14, "2": 15, "B": 16, "C": 17}
    mp2 := map[int]string{3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10", 11: "J", 12: "Q", 13: "K", 14: "A", 15: "2", 16: "B", 17: "C"}
    for i := 0; i < len(strs1); i++ {
        remain[mp[strs1[i]]]--
    }
    for i := 0; i < len(strs2); i++ {
        remain[mp[strs2[i]]]--
    }
    fmt.Println(remain)

    // 开始找剩下的牌最长的顺子
    var maxSub, sub string
    var maxCount, count int
    for i := 0; i < len(remain); i++ {
        if remain[i] > 0 {
            sub = sub + mp2[i] + "-"
            count++
        } else if len(maxSub) < len(sub) {
            maxSub = sub
            maxCount = count
            sub = ""
            count = 0
            fmt.Println(sub)
        }
    }
    if maxCount < 5 {
        fmt.Println("NO-CHAIN")
        return
    }
    fmt.Println(maxSub[:len(maxSub)-1])
    fmt.Println(maxCount)
}
