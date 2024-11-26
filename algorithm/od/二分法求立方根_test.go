package od

import (
    "fmt"
    "testing"
)

func TestGaskl(t *testing.T) {
    var count float64
    fmt.Scan(&count)

    if count > 0 {
        fmt.Printf("%0.1f\n", gen(count))
    } else {
        fmt.Printf("%0.1f\n", -1.0*gen(-1.0*count))
    }

}
func gen(n float64) float64 {
    var j, s float64
    min, max := 0.0, n
    if n < 1 { // 因为n<1时，n是要小于解的，而循环的范围是0-n的话就永远找不到解
        max = 1
    }
    for max-min > 0.00000001 {
        j = (max + min) / 2
        s = j * j * j
        if s > n {
            max = j
        } else {
            min = j
        }
    }
    return j
}
