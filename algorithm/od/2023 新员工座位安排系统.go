package od

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
/**
工位由序列F1,F2…Fn组成，Fi值为0、1或2。其中0代表空置，1代表有人，2代表障碍物。

1、某一空位的友好度为左右连续老员工数之和，

2、为方便新员工学习求助，优先安排友好度高的空位，

给出工位序列，求所有空位中友好度的最大值。


*/
//双指针
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()

    trim := strings.Trim(text, " ")
    split := strings.Split(trim, " ")
    var max, index int
    for i := 0; i < len(split); i++ {
        if split[i] == "0" {
            a, b, count := i-1, i+1, 0
            for a >= 0 && split[a] == "1" {
                count++
                a--
            }
            for b < len(split) && split[b] == "1" {
                count++
                b++
            }
            if max < count {
                max = count
                index = i
            }
        }
    }
    fmt.Println("友好度最大值：", max)
    fmt.Println("位置下标：", index)

}

