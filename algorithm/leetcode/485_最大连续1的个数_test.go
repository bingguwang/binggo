package leetCode

import (
    "fmt"
    "testing"
)

/**
1 1 0 1 1 1
i

*/
func TestJKan(t *testing.T) {
    ones := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
    //ones := findMaxConsecutiveOnes([]int{0})
    fmt.Println(ones)
}
func findMaxConsecutiveOnes(nums []int) (maxCnt int) {
    cnt := 0
    for _, v := range nums {
        if v == 1 {
            cnt++
        } else {
            maxCnt = max(maxCnt, cnt)
            cnt = 0
        }
    }
    maxCnt = max(maxCnt, cnt)
    return

}
