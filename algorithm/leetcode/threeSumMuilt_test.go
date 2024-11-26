package leetCode

import (
    "fmt"
    "sort"
    "testing"
)

func TestKls(t *testing.T) {
    multi := threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8)
    fmt.Println(multi)
}
func threeSumMulti(arr []int, target int) int {
    var ret int
    mp := map[int]int{}
    for _, v := range arr {
        mp[v]++
    }
    fmt.Println(mp)
    var uniquNums []int
    for k, _ := range mp {
        uniquNums = append(uniquNums, k)
    }
    sort.Ints(uniquNums)
    fmt.Println(uniquNums)
    for i := 0; i < len(uniquNums); i++ {
        ni := mp[uniquNums[i]]
        if uniquNums[i]*3 == target && mp[uniquNums[i]] >= 3 {
            ret += ni * (ni - 1) * (ni - 2) / 6
        }
        for j := i + 1; j < len(uniquNums); j++ {
            nj := mp[uniquNums[j]]
            if uniquNums[i]*2+uniquNums[j] == target && mp[uniquNums[i]] >= 2 && mp[uniquNums[j]] >= 1 {
                ret += ni * (ni - 1) / 2 * nj
            }
            if uniquNums[j]*2+uniquNums[i] == target && mp[uniquNums[j]] >= 2 && mp[uniquNums[i]] >= 1 {
                ret += nj * (nj - 1) / 2 * ni
            }
            c := target - uniquNums[i] - uniquNums[j]
            if c > uniquNums[j] && mp[c] >= 1 {
                ret += ni * nj * mp[c]
            }

        }
    }
    return ret % (10e9 + 7)
}



