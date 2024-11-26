package Backtracking

import (
    "fmt"
    "testing"
)

func TestBT(t *testing.T)  {
    items := []Item{
        {weight: 5, value: 10},
        {weight: 3, value: 8},
        {weight: 4, value: 7},
        {weight: 2, value: 6},
    }
    maxWeight := 9

    fmt.Println("最大价值:", knapsack(items, maxWeight))


}
type Item struct {
    weight int
    value  int
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func knapsack(items []Item, maxWeight int) int {
    return knapsackHelper(items, maxWeight, 0, 0)
}

// 选择index节点
func knapsackHelper(items []Item, maxWeight int, index int, value int) int {
    if index >= len(items) {
        return value
    }

    // 当前index节点的重量超过最大容量，则不能选当前节点
    if items[index].weight > maxWeight {
        fmt.Println("不能选择节点: ", items[index])
        return knapsackHelper(items, maxWeight, index+1, value)
    }

    return max(
        // 不选择当前节点
        knapsackHelper(items, maxWeight, index+1, value),
        // 选择当前节点
        knapsackHelper(items, maxWeight-items[index].weight, index+1, value+items[index].value),
    )
}
