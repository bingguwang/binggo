package od

import (
    "fmt"
    "math"
    "testing"
)

/**
  把一个16进制数转为十进制
*/

func TestGasdlq(t *testing.T) {
    change := Change("0xAA")
    fmt.Println(change)
}

func Change(a string) int {
    var res int
    var str string
    for i := 2; i < len(a); i++ {
        str += Ex(a[i])
    }
    fmt.Println(str)
    for i := len(str) - 1; i >= 0; i-- {
        val := int(byte(str[i]) - byte('0'))
        if val == 0 {
            continue
        }
        //fmt.Printf("%v ", val)
        m := len(str) - 1 - i
        //fmt.Printf("%v ", m)
        res += val * int(math.Pow(2, float64(m)))

    }

    return res
}
func Ex(a byte) string {
    switch string(a) {
    case "A":
        return "1010"
    case "B":
        return "1011"
    case "C":
        return "1100"
    case "D":
        return "1101"
    case "E":
        return "1110"
    case "F":
        return "1111"
    case "0":
        return "0000"
    case "1":
        return "0001"
    case "2":
        return "0010"
    case "3":
        return "0011"
    case "4":
        return "0100"
    case "5":
        return "0101"
    case "6":
        return "0110"
    case "7":
        return "0111"
    case "8":
        return "1000"
    case "9":
        return "1001"
    }
    return ""
}
/**
    还有更合适的方式，直接用fmt包就够了
 */
func TestRgasjq(t *testing.T)  {
    var str int
    fmt.Scanf("0x%x", &str) // 输入格式是0xAA这样的
    fmt.Println(str)

}