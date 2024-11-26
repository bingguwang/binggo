package od

import (
    "fmt"
    "sort"
)

/**
 给你一个整数M和数组N,N中的元素为  连续整数，要求根据N中的元素组装成新的数组R，组装规则：

1.R中元素总和加起来等于M
2.R中的元素可以从N中重复选取
3.R中的元素最多只能有1个不在N中，且比N中的数字都要小（不能为负数）

请输出：数组R一共有多少组装办法

// 思路：回溯递归

*/
func main() {
    ways := countAssembleWays(5, []int{2, 3})
    fmt.Println(ways)

}
func countAssembleWays(M int, N []int) int {
    sort.Ints(N)
    min := N[0]
    max := N[len(N)-1]
    res := 0

    var assembleHelper func(int, int)
    assembleHelper = func(v, sum int) { // 传入v当前选择的值， 当前还需要sum
        if min > sum {                  // 当前还需要的值比min还小
            res++ // N里的最小值都比剩下需要的值大，则可以选择一个N外的数
            return
        }
        if max < v { // 当前选择的值不能超过N里的最大值
            return
        }

        i := 0
        for {
            // 因为数组中的数字是连续的，所以只需要+1
            assembleHelper(v+1, sum-i*v)
            i++               // 看取几个该值合适
            if sum-i*v <= 0 { // 超过了
                if sum-i*v == 0 { // 有个解了
                    res++
                }
                break
            }
        }
    }
    assembleHelper(min, M)
    return res
}
