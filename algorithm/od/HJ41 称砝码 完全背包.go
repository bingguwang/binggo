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
    fmt.Println("maxsum: ", sum)
    weigh(n, weights, nums, sum)
}
func weigh(n int, m []int, x []int, maxsum int) int {
    dp := make([]bool, maxsum+1)
    // 初始化
    dp[0] = true
    dp[maxsum] = true
    for i := 0; i < n; i++ { // 物品种类遍历， 砝码种类遍历
        for j := maxsum; j >= m[i]; j-- { //  for j := maxsum; j >  0; j-- 的优化而已
            for k := 1; k <= x[i]; k++ { // 选取k个第i种砝码
                if j-k*m[i] >= 0 { // 不能选
                    dp[j] = dp[j] || dp[j-k*m[i]]
                } else {
                    dp[j] = dp[j]
                }
            }
        }
    }
    count := 0
    for _, v := range dp {
        if v {
            count++
        }
    }
    fmt.Println(dp)
    fmt.Println(count)
    return count
}
