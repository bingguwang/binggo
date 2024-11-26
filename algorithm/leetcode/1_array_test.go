package leetCode

import (
    "fmt"
    "testing"
)

/**
tar 6
2 3 3 4 1

在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
*/
func twoSums(nums []int, target int) []int {
    mp := make(map[int]int)

    for i, num := range nums {
        if v, ok := mp[num]; ok {
            return []int{v, i}
        }
        mp[target-num] = i
    }
    return nil
}
func F(a []int, tar int) []int {
    mp := make(map[int]int)
    for i, v := range a {
        if idx, ok := mp[v]; ok && idx != i {
            return []int{i, idx}
        }
        mp[tar-v] = i
    }
    return nil
}

func TestS(t *testing.T) {
    sums := twoSums([]int{2, 3, 3, 3, 1}, 6)
    fmt.Println(sums)
}
