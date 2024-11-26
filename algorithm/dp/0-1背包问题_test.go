package dp

import (
	"fmt"
	"testing"
)

// dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]] + v[i])
//表示i个物品放入容量为j的背包的最大价值

func TestRS(t *testing.T) {
	items := []Item{
		{weight: 5, value: 10},
		{weight: 3, value: 8},
		{weight: 4, value: 7},
		{weight: 2, value: 6},
	}
	Dp(items, 9)
}

func Dp(items []Item, maxCap int) {
	dp := make([][]int, len(items))
	for i := 0; i < len(dp); i++ {
		p := make([]int, maxCap+1) // 不加1会少一列结果
		dp[i] = p
	}
	fmt.Println(dp)

	// 边界状态
	for i := 0; i < len(items); i++ {
		dp[i][0] = 0
	}
	for i := 0; i < maxCap+1; i++ {
		dp[0][i] = 0
	}
	for _, v := range dp {
		fmt.Println(v)
	}

	//for i := 1; i < len(items); i++ {
	//    for j := 1; j < maxCap+1; j++ {
	//        // 物品重量超过容量，则不能选物品,只有不选一种决策
	//        if items[i].weight > j {
	//            dp[i][j] = dp[i-1][j]
	//        } else { // 可以选物品，选或者不选
	//            dp[i][j] = max22(dp[i-1][j], dp[i-1][j-items[i].weight]+items[i].value)
	//        }
	//    }
	//}

	for i := 1; i < len(items); i++ {
		for j := maxCap; j >= 0; j-- {
			// 物品重量超过容量，则不能选物品,只有不选一种决策
			if items[i].weight > j {
				dp[i][j] = dp[i-1][j]
			} else { // 可以选物品，选或者不选
				dp[i][j] = max22(dp[i-1][j], dp[i-1][j-items[i].weight]+items[i].value)
			}
		}
	}

	fmt.Println("----------")
	for _, v := range dp {
		fmt.Println(v)
	}
	fmt.Println(dp[len(items)-1][maxCap])
}

type Item struct {
	weight int
	value  int
}
