package leetCode

import (
    "fmt"
    "sort"
    "testing"
)

/**

给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合

*/
func TestJsdk(t *testing.T) {
    //sums := threeSums([]int{-1, 0, 1, 2, -1, -4})
    sums := threeSums2([]int{-1, 0, 1, 1, 2, -1, -4})
    //sums := threeSums2([]int{0, 0, 0, 0})
    //sums := threeSums([]int{-2, 0, 1, 1, 2})
    fmt.Println(sums)
}

// 排序后 双指针
func threeSums(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    fmt.Println(nums)
    for idx := 1; idx < len(nums); idx++ {
        start, end := 0, len(nums)-1
        if idx > 1 && nums[idx-1] == nums[idx] {
            start = idx - 1
        }
        for start < idx && end > idx {
            if start > 0 && nums[start-1] == nums[start] {
                start++
                continue
            }
            if end < len(nums)-1 && nums[end] == nums[end+1] {
                end--
                continue
            }

            sum := nums[start] + nums[end] + nums[idx]
            if sum == 0 {
                res = append(res, []int{nums[start], nums[idx], nums[end]})
                end--
                start++
            } else if sum < 0 {
                start++
            } else {
                end--
            }
        }
    }
    return res
}

// 排序后暴力法
func threeSums2(nums []int) [][]int {
    var res [][]int
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    var uniqueArr []int
    for k, _ := range mp {
        uniqueArr = append(uniqueArr, k)
    }
    sort.Ints(uniqueArr)

    for i := 0; i < len(uniqueArr); i++ {
        if uniqueArr[i]*3 == 0 && mp[uniqueArr[i]] >= 3 {
            res = append(res, []int{uniqueArr[i], uniqueArr[i], uniqueArr[i]})
        }

        for j := i + 1; j < len(uniqueArr); j++ {
            if uniqueArr[i]*2+uniqueArr[j] == 0 && mp[uniqueArr[i]] >= 2 && mp[uniqueArr[j]] >= 1 {
                res = append(res, []int{uniqueArr[i], uniqueArr[i], uniqueArr[j]})
            }
            if uniqueArr[j]*2+uniqueArr[i] == 0 && mp[uniqueArr[j]] >= 2 && mp[uniqueArr[i]] >= 1 {
                res = append(res, []int{uniqueArr[i], uniqueArr[j], uniqueArr[j]})
            }

            c := 0 - uniqueArr[i] - uniqueArr[j]
            if c > uniqueArr[j] && mp[c] > 0 { // c要是存在的
                res = append(res, []int{uniqueArr[i], uniqueArr[j], c})
            }
        }
    }
    return res

}
