package first

import (
    "fmt"
    "testing"
)

/**

给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

类似单词拆分的那题

*/
func TestMksd(t *testing.T) {
    //a := []int{2, 3, 1, 1, 4}
    a := []int{0}
    number := jump(a)
    fmt.Println(number)
}
func jump(nums []int) int {
    dp := make([]int, len(nums))
    for i := 0; i < len(dp); i++ {
        dp[i] = 10000
    }
    // dp[i]到达第i个元素最小跳跃次数, 也就是0到i-1这段
    dp[0] = 0
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if j+nums[j] >= i { // 在位置j可以够得到i
                dp[i] = min(dp[i], dp[j]+1)
            }
        }
    }
    fmt.Println(dp)
    return dp[len(nums)-1]
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
