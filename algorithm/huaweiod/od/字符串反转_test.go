package od

import (
    "fmt"
    "testing"
)

func TestGHlsa(t *testing.T) {
    str := reverseStr([]byte("abcde"))
    fmt.Println(str)
}
func reverseStr(str []byte) string {
    i, j := 0, len(str)-1
    for i < j {
        str[i], str[j] = str[j],str[i]
        j--
        i++
    }
    return string(str)
}
