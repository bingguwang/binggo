package od

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


/**
题目描述：
   某云短信运营商，为庆祝国庆，推出充值优惠活动。现在给出客户预算，和优惠售价序列，求最多可获得的短信总条数。
输入描述：第一行客户预算M，其中0 <= M <= 10^6
         第二行给出售价表，P1,P2,...Pn，其中 1 <= n <= 100，
         Pi为充值i元获得的短信条数。1 <= Pi <= 1000，1 <= n <= 100
输出描述：
最多获得的短信条数。

示例

示例1
输入输出示例仅供调试，后台判题数据一般不包含示例
输入：
6
10 20 30 40 60
输出：
70
说明：分两次充值最优，1元、5元各一次。总条数 10 + 60 = 70

示例2
输入：
15
10 20 30 40 60 60 70 80 90 150
输出：
210
6
10 20 30 55 60

动态规划
*/

func main() {
    var budget int
    fmt.Scan(&budget)

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    split := strings.Split(text, " ")
    var smsCounts []int
    for _, s := range split {
        atoi, _ := strconv.Atoi(s)
        smsCounts = append(smsCounts, atoi)
    }

    n := len(smsCounts)
    // 初始化动态规划表，dp[i][j]表示前i个售价，预算为j时最多可获得的短信条数
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, budget+1)
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= budget; j++ {
            k := 0
            nowsum := 0
            for {
                nowsum += k * i
                if nowsum > j {
                    break
                }
                dp[i][j] = max(dp[i][j], dp[i-1][j-k*i]+k*smsCounts[i-1])
                k++
            }
        }
    }

    fmt.Println(dp[n][budget])
}
