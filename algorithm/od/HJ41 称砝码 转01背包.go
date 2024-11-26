package od

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    nums := make([]int, n)
    weights := make([]int, n)
    for i := 0; i < len(weights); i++ {
        fmt.Scan(&weights[i])
    }
    sum := 0
    for i := 0; i < len(nums); i++ {
        fmt.Scan(&nums[i])
        sum += weights[i] * nums[i]
    }

    // 存储全部类的砝码，把每个砝码作为一种商品，然后使用0-1背包问题来使用DP
    var wts []int
    for i := 0; i < n; i++ {
        for j := 0; j < nums[i]; j++ {
            wts = append(wts, weights[i])
        }
    }
    // 至此，可以转为01背包问题

    // dp[i][j]就是背包问题里的，表示前i种商品前的商品凑出j的可行性
    dp := make([][]bool, len(wts)+1)
    for i := 0; i < len(wts)+1; i++ {
        dp[i] = make([]bool, sum+1)
    }

    mp := make(map[int]bool)
    for i := 0; i < len(wts)+1; i++ {
        dp[i][0] = true
        mp[0] = true
    }
    for i := 1; i < len(wts)+1; i++ {
        for j := 1; j < sum+1; j++ {
            if wts[i-1] > j { // wts[i-1]表示当前要选的物品
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-wts[i-1]]
            }
            if dp[i][j] {
                mp[j] = true
            }
        }
    }
    count := len(mp)
    fmt.Println(count)
}
