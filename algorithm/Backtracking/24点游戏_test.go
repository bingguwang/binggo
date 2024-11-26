package Backtracking

import (
    "fmt"
    "math"
    "testing"
)
/**
思路是：
游戏的第一步是挑出两个数，算出一个新数替代这两个数。
然后，在三个数中玩 24 点，再挑出两个数，算出一个数替代它们。
然后，在两个数中玩 24 点……
最后只有一个数了， 我们就只要判断这个数是不是等于24就行了


因为是可以使用括号的，所以上面的思路是可行的

考察的是dfs,回溯思想，一般使用递归实现
 */



func dfs(nums []float64) bool {
    if len(nums) == 1 { // 数组只有1个数了
        return math.Abs(nums[0]-24) < 1e-9 // 是24
    }
    flag := false
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ { // 2层for循环枚举出了所有:挑出两个数  的可能
            n1, n2 := nums[i], nums[j]
            newNums := make([]float64, 0, len(nums)) // numNums保存的是两数计算结果替换两数后，新的数组
            for k := 0; k < len(nums); k++ {
                if k != i && k != j {
                    newNums = append(newNums, nums[k]) // 先把没被选中的数加到numNums里
                }
            }

            flag = flag || dfs(append(newNums, n1+n2))
            flag = flag || dfs(append(newNums, n1-n2))
            flag = flag || dfs(append(newNums, n2-n1))
            flag = flag || dfs(append(newNums, n1*n2))
            // 除法防止除0
            if n1 != 0 {
                flag = flag || dfs(append(newNums, n2/n1))
            }
            if n2 != 0 {
                flag = flag || dfs(append(newNums, n1/n2))
            }
            if flag {
                return true
            }
        }
    }
    return false
}

func TestSD(t *testing.T)  {
    point24 := dfs([]float64{1, 3, 1, 5})
    fmt.Println(point24)
}