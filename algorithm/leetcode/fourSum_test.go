package leetCode

import (
    "fmt"
    "sort"
    "testing"
)

func TestAS(t *testing.T) {
    //sum := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
    //sum := fourSum([]int{2, 2, 2, 2, 2}, 8)
    sum := fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
    fmt.Println(sum)
}
func fourSum(nums []int, target int) [][]int {
    var ret [][]int
    sort.Ints(nums)
    n := len(nums)
    for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ { // i++只会越来越大，四数之和如果大于tar,后面没必要遍历
        if i > 0 && nums[i-1] == nums[i] {
            continue
        }
        for j := i + 1; j < n-2  && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
            if nums[j-1] == nums[j] && j-1 > i {
                continue
            }
            for l, r := j+1, n-1; l < r; {
                if nums[l] == nums[l-1] && l-1 > j {
                    l++
                    continue
                }
                sum := nums[i] + nums[j] + nums[l] + nums[r]
                if sum == target {
                    ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
                    l++
                    r--
                } else if sum < target {
                    l++
                } else {
                    r--
                }

            }
        }
    }
    return ret
}

/**
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
nums[a] + nums[b] + nums[c] + nums[d] == target

nums = [1,0,-1,0,-2,2], target = 0
[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
[-2,-1,-1,1,1,2,2]

2 2 2 2 2
-2 -1 -1 1 1 2 2
i
    j
         l
               r

*/
