package leetCode

import (
    "fmt"
    "math"
    "sort"
    "testing"
)

func TestASA(T *testing.T) {
    //sum := threeSumClosest([]int{-1, 2, 1, -4}, 1)
    //sum := threeSumClosest([]int{1, 1, 1, 0}, -100)
    sum := threeSumClosest([]int{1,2,4,8,16,32,64,128}, 82)
    fmt.Println(sum)
}

/**
1,2,4,8,16,32,64,128
  i
        l
               r
*/
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    fmt.Println(nums)
    dif, ret := math.MaxInt, 0
    for i := 0; i < len(nums); i++ {

        for l, r := i+1, len(nums)-1; l < r; {
            sum := nums[i] + nums[l] + nums[r]
            x := abs(sum, target)
            if x < dif {
                dif = x
                ret = sum
            }
            if sum == target {
                return target
            } else if sum < target {
                l++
            } else {
                r--
            }
        }

    }
    return ret

}
func abs(a, b int) int {
    if a > b {
        return a - b
    } else {
        return b - a
    }
}

/**
nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
-4 -1 1 2
i
    l
        r

*/
