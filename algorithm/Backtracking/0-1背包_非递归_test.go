package Backtracking

import (
    "fmt"
    "testing"
)

type State struct {
    index  int
    weight int
    value  int
}

func knapsack2(items []Item, maxWeight int) int {
    maxVal := 0
    stack := []State{{index: 0, weight: 0, value: 0}}
    for len(stack) > 0 {
        fmt.Println("栈内情况:  ", stack)

        // 出栈
        s := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        fmt.Println("出栈:", s)

        if s.index >= len(items) { // 节点都在栈里了，此时必须要出栈
            fmt.Println("节点都在站内了，也就是到了搜索树的叶子节点了！，此时必须直接出栈")
            maxVal = max(maxVal, s.value)
            continue
        }

        // 出栈节点重量 + 当前节点重量 < 总容量
        fmt.Println("s.weight :", s.weight)
        fmt.Println("items[s.index].weight 下一个重量 :", items[s.index].weight)
        if s.weight+items[s.index].weight <= maxWeight { // 可以放入，就需要放入两个节点。
            p := State{s.index + 1, s.weight + items[s.index].weight, s.value + items[s.index].value}
            fmt.Println("入栈 ", p)
            stack = append(stack, p)
        }
        // 不能放入我们就只放入一个节点
        p2 := State{s.index + 1, s.weight, s.value}
        fmt.Println("入栈 ", p2)
        stack = append(stack, p2)
    }
    return maxVal
}

func TestR(t *testing.T) {
    items := []Item{
        {weight: 5, value: 10},
        {weight: 3, value: 8},
        {weight: 4, value: 7},
        {weight: 2, value: 6},
    }

    maxWeight := 9

    fmt.Println("最大价值:", knapsack2(items, maxWeight))
}
