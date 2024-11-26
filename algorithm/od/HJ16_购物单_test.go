package od

import (
    "fmt"
    "testing"
)


type Item struct {
    price  int
    weight int
    isMain bool
    a1     int
    a2     int
    p      int
}

func TestHgs(t *testing.T) {
    var N, m int
    fmt.Scan(&N, &m)
    goods := make([]Item, m)
    for i := 0; i < m; i++ {
        goods[i] = Item{a1: -1, a2: -1}
    }
    for i := 0; i < m; i++ {
        var v, p, q int
        fmt.Scan(&v, &p, &q)
        goods[i].price = v
        goods[i].weight = p
        goods[i].p = p * v
        if q == 0 {
            goods[i].isMain = true
        } else if goods[q-1].a1 == -1 {
            goods[q-1].a1 = i
        } else {
            goods[q-1].a2 = i
        }
    }
    // dp[i][j]表示前i个物品里, 总金额不超过j,此时的最大满意度
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, N+1)
    }

    for i := 1; i <= m; i++ {
        for j := 0; j <= N; j++ {
            dp[i][j] = dp[i-1][j]
            if !goods[i-1].isMain {
                continue
            }
            if j >= goods[i-1].price {
                dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i-1].price]+goods[i-1].p)
            }
            if goods[i-1].a1 != -1 && j >= goods[i-1].price+goods[goods[i-1].a1].price { // 第i个是主物品，且有一个附件1,且在奖金范围内
                dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i-1].price-goods[goods[i-1].a1].price]+goods[i-1].p+goods[goods[i-1].a1].p)
            }
            if goods[i-1].a2 != -1 && j >= goods[i-1].price+goods[goods[i-1].a2].price { // 第i个是主物品，且有一个附件2,且在奖金范围内
                dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i-1].price-goods[goods[i-1].a2].price]+goods[i-1].p+goods[goods[i-1].a2].p)
            }
            if goods[i-1].a1 != -1 && goods[i-1].a2 != -1 && j >= goods[i-1].price+goods[goods[i-1].a1].price+goods[goods[i-1].a2].price { // 第i个是主物品，且有2个附件 ,且在奖金范围内
                dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i-1].price-goods[goods[i-1].a1].price-goods[goods[i-1].a2].price]+goods[i-1].p+goods[goods[i-1].a1].p+goods[goods[i-1].a2].p)
            }
        }
    }
    fmt.Println(dp[m][N])
}
