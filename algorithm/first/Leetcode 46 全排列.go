package first

import (
    "fmt"
    "testing"
)

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

有点类似牛客的火车进站问题
用回溯

*/
func TestFGJADQ(t *testing.T) {
    duplicate := permute([]int{1, 3, 2})
    fmt.Println(duplicate)
}
func permute(nums []int) [][]int {
    var res [][]int
    n := len(nums)
    visted := make([]bool, n)
    var dfs func([]int)
    dfs = func(ints []int) {
        if len(ints) == n {
            //得到了一个解
            tmp := make([]int, n)
            copy(tmp, ints)
            res = append(res, tmp) // 不能直接append ints，因为ints其实是指向它的底层数组的指针，如果后续修改了ints，就会影响已经添加到结果中的排列，导致错误的结果。res是二维的呦，可以懂了吧
            return
        }
        for i := 0; i < n; i++ {
            if !visted[i] { // 如果没有访问过
                visted[i] = true
                ints = append(ints, nums[i])
                dfs(ints)
                // 回溯,恢复
                ints = ints[:len(ints)-1]
                visted[i] = false
            }
        }

    }
    dfs([]int{})
    return res
}
