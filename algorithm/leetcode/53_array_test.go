package leetCode

import (
    "fmt"
    "testing"
)

/**

找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。
*/
func TestSk(t *testing.T) {
    //array := maxSubArrays([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
    //array := maxSubArrays([]int{5, 4, -1, 7, 8})
    array := maxSubArrays([]int{-1, -2})
    fmt.Println(array)
}

// DP : di的值根据d(i-1)正负来决定，di表示必须算i在内的最大
func maxSubArray(nums []int) int {
    dp := make([]int, len(nums))
    var res int
    dp[0] = nums[0]
    res = dp[0]
    for i := 1; i < len(nums); i++ {
        if dp[i-1] > 0 {
            dp[i] = dp[i-1] + nums[i]
        } else {
            dp[i] = nums[i]
        }
        res = max(dp[i], res)
    }
    return res
}

// 一个指针右移，把移动的值都加上，为负了就把和置为0，这是一般的思考其实
func maxSubArrays(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    res, sum, p := nums[0], 0, 0
    for p < len(nums) {
        sum += nums[p]
        if sum > res {
            res = sum
        }
        if sum <= 0 {
            sum = 0
        }
        p++
    }
    return res
}
