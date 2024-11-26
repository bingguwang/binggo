package leetCode

import (
    "fmt"
    "testing"
)

/**
给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。
任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。
然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的 最短时间 。

【解决】:桶思想，每个桶长度为n+1, n冷却时间
*/
func TestGhakw(t *testing.T) {
    interval := leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
    fmt.Println(interval)

}
func leastInterval(tasks []byte, n int) int {
    arr := make([]int, 26)
    for _, v := range tasks {
        arr[v-'A']++
    }
    fmt.Println(arr)

    m := 0 // 最大任务数
    c := 0 // 最大任务数的数量

    for i := 0; i < len(arr); i++ {
        if m < arr[i] {
            m = arr[i]
            c = 1
        } else if m == arr[i] {
            c++
        }
    }
    if len(tasks) > (m-1)*(n+1)+c {
        return len(tasks)
    }
    return (m-1)*(n+1) + c
}
