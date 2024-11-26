package doublepointer

import (
    "fmt"
    "testing"
)

func TestAs(T *testing.T) {
    youxu := twoSumYouxu([]int{2, 7, 11, 15}, 9)
    //youxu := twoSumYouxu([]int{2, 3, 4}, 6)
    //youxu := twoSumYouxu([]int{-1, 0}, -1)
    fmt.Println(youxu)
}
func twoSumYouxu2(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        for j := i + 1; j < len(numbers); j++ {
            if target == numbers[i]+numbers[j] {
                return []int{i + 1, j + 1}
            }
        }
    }
    return nil
}
func twoSumYouxu(numbers []int, target int) []int {
    l, r := 0, len(numbers)-1
    for l < r {
        sum := numbers[l] + numbers[r]
        if sum == target {
            return []int{l + 1, r + 1}
        } else if sum < target {
            l++
        } else {
            r--
        }
    }
    return nil
}
