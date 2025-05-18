package od

import (
    "fmt"
    "strings"
    "testing"
)

func TestJks(t *testing.T) {
    a := "5.2"
    b := "5.02"
    compodsa := Compodsa(a, b)
    fmt.Println(compodsa)
}
func Compodsa(a, b string) int {
    a = strings.TrimSpace(a)
    b = strings.TrimSpace(b)
    as := strings.Split(a, ".")
    bs := strings.Split(b, ".")
    fmt.Println(as)
    fmt.Println(bs)
    i := 0
    for i < len(as) && i < len(bs) {
        x := strings.TrimLeft(as[i], "0")
        y := strings.TrimLeft(bs[i], "0")
        if x < y {
            return -1
        }
        if x > y {
            return 1
        }
        i++
    }
    if len(as) == len(bs) {
        return 0
    }

    if i >= len(as) {
        return 1
    }
    return -1
}
