package first

import (
    "fmt"
    "sort"
    "testing"
)
/**
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

 */
func TestFJHADMB(t *testing.T) {
    //duplicate := findDuplicate([]int{1, 3, 4, 2, 2})
    duplicate := findDuplicate([]int{2, 2, 2, 2, 2})
    fmt.Println(duplicate)
}

// 暴力一点。排序
func findDuplicate(nums []int) int {
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if i-1 >= 0 && nums[i] == nums[i-1] {
            return nums[i]
        }
    }
    return -1
}

// 二分法， 抽屉原理
func findDuplicate2(nums []int) int {
    left, right := 1, len(nums)-1 // 长度为 n + 1 的数组，数值在 1 到 n 之间。因此长度为 len = n + 1，n = len - 1，搜索范围在 1 到 len - 1 之间；
    for left < right {
        mid := left + (right-left)/2 // 防止left+right)/2 时left+right溢出
        cnt := 0
        for _, num := range nums { // 统计 mid以及mid左边有多少个数
            if num <= mid {
                cnt++
            }
        }
        if cnt > mid { // mid左边数多
            right = mid
        } else { // mid右边数多
            left = mid + 1
        }
    }
    return left
}
