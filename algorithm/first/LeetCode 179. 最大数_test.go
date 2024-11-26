package first

import (
    "fmt"
    "sort"
    "strconv"
    "testing"
)

func TestGAJDB(t *testing.T) {
    a := []int{3, 30, 34, 5, 9}
    number := largestNumber(a)
    fmt.Println(number)
}
func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        x, y := nums[i], nums[j]
        sx, sy := 10, 10
        for sx <= x {
            sx *= 10
        }
        for sy <= y {
            sy *= 10
        }
        // 转为同位后比较
        return sy*x+y > sx*y+x
    })
    fmt.Println(nums)
    var res string
    for i := 0; i < len(nums); i++ {
        res += strconv.FormatInt(int64(nums[i]), 10)
    }
    return res
}
