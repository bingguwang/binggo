package leetCode

import (
    "fmt"
    "testing"
)

/**
  合并有序数组
*/

func TestJksad(t *testing.T) {
    merge([]int{1, 6, 7}, 3, []int{2, 4, 9, 11}, 4)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
    i, j := 0, 0
    var sorted []int
    for {
        if i == m {
            sorted = append(sorted, nums2[j:]...)
            break
        }
        if j == n {
            sorted = append(sorted, nums1[i:]...)
            break
        }
        if nums1[i] < nums2[j] {
            sorted = append(sorted, nums1[i])
            i++
        } else {
            sorted = append(sorted, nums2[j])
            j++
        }
    }
    fmt.Println(sorted)
    copy(nums1, sorted)
}

/**
  1 6 7
  2 4 9 11

*/
