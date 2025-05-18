package od

import (
    "fmt"
    "sort"
)

/**

时间限制：1s 空间限制：256MB 限定语言：不限

题目描述：

小明有n块木板，第i（1＜=i＜=n）块木板的长度为ai。

小明买了一块长度为m的木料，这块木料可以切割成任意块，拼接到已有的木板上，用来加长木板。

小明想让最短的木板尽量长。

请问小明加长木板后，最短木板的长度最大可以为多少？
*/

// 使用排序来实现
func main() {
    var n, m int
    fmt.Scan(&n, &m)
    a := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }

    for i := 0; i < m; i++ {
        sort.Ints(a)
        a[0]++
    }
    sort.Ints(a)
    fmt.Println(a[0])
}
