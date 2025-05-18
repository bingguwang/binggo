package od

import (
    "fmt"
    "sort"
    "testing"
)
/**
给定一个正整数N代表火车数量，0<N<10，接下来输入火车入站的序列，一共N辆火车，每辆火车以数字1-9编号，火车站只有一个方向进出，同时停靠在火车站的列车中，只有后进站的出站了，先进站的才能出站。
要求输出所有火车出站的方案，以字典序排序输出。
数据范围：1\le n\le 10\1≤n≤10
进阶：时间复杂度：O(n!)\O(n!) ，空间复杂度：O(n)\O(n)
输入描述：
第一行输入一个正整数N（0 < N <= 10），第二行包括N个正整数，范围为1到10。

输出描述：
输出以字典序从小到大排序的火车出站序列号，每个编号以空格隔开，每个输出序列换行，具体见sample。
 */
func TestFSAD(t *testing.T) {
    var train [13]int
    var n int // 火车数
    var paths []string
    var path []int
    var stack []int

    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&train[i])
    }

    var dfs func(int)
    dfs = func(i int) {     // i 是要决定是否入栈的火车编号,i是入栈序列的编号
        if len(path) == n { // 一个方案完成
            var s string
            for k := 0; k < n; k++ {
                s += fmt.Sprintf("%d ", path[k])
            }
            paths = append(paths, s)
            return
        }
        if len(stack) != 0 { // 选择栈内火车出栈
            // 记录栈顶。栈顶出栈
            path = append(path, stack[len(stack)-1])
            stack = stack[:len(stack)-1]

            dfs(i) // 因为选择的是火车出栈，i没有进入栈内，所以子递归还是i
            // 回溯，撤销"选择栈内火车出栈"造成的结果
            stack = append(stack, path[len(path)-1])
            path = path[:len(path)-1]
        }
        if i < n { //  选择让火车入栈
            stack = append(stack, train[i])
            dfs(i + 1)                   // 子递归去让下一个火车入栈
            stack = stack[:len(stack)-1] // 回溯，撤销掉“ 选择让火车入栈”导致的结果
        }
    }
    // 执行
    dfs(0)

    sort.Strings(paths)
    for _, s := range paths {
        fmt.Println(s)
    }
}
