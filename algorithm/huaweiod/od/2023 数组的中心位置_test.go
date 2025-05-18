package od

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

/**
给你一个整数数组nums,请计算数组的中心位置 。数组中心位置是数组的一个下标，其左侧所有元素相乘的积等于右侧所有元素相乘的积。
数组第一个元素的左侧积为1，最后一个元素的右侧积为1
如果数组有多个中心位置，应该返回最靠近左边的那一个。如果数组不存在中心位置，返回 -1 。
*/
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    text = strings.Trim(text, " ")
    split := strings.Split(text, " ")

    a := make([]int, len(split))
    for i := 0; i < len(split); i++ {
        atoi, _ := strconv.Atoi(split[i])
        a[i] = atoi
    }

    for i := 0; i < len(a); i++ {
        k, p := i-1, i+1
        leftmul, rightmul := 1, 1
        for k >= 0 {
            if k == 0 { // 两端特殊处理
                leftmul *= 1
            }
            leftmul = leftmul * a[k]
            k--
        }
        for p < len(a) {
            if p == len(a) {
                rightmul *= 1
            }
            rightmul = rightmul * a[p]
            p++
        }
        fmt.Println(leftmul, " ", rightmul)
        if leftmul == rightmul {
            fmt.Println(i)
            break
        }
    }
}
