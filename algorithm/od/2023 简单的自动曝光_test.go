package od

import (
    "fmt"
    "math"
)

/**
一个图像有n个像素点，存储在一个长度为n的数组img里，每个像素点的取值范围[0,255]的正整数请你给图像每个像素点值加上一个整数k(可以是负数)，得到新图newimg，使得新图newimg的所有像素平均值最接近128请输出这个整数k。
输入描述
n个整数，中间用空格分开
输出描述
个整数k
备注
1 <= n <= 100
如有多个整数k都满足，输出小的那个k;
新图的像素值会自动截取到10.2551范围。当新像素值<0，其值会更改为0;当新像素值>255，其值会更改为255:
例如newlmg=”-1 -2 256”,会自动更改为”0 0 255"
请你用go实现
*/

func main() {
    // 读取输入
    var n int
    fmt.Scan(&n)
    img := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&img[i])
    }

    // 计算原图像素平均值
    avg := 0
    for _, val := range img {
        avg += val
    }
    avg /= n

    // 找到最接近128的k
    diff := math.MaxInt32
    var k int
    //因为像素值的取值范围是 [0, 255]，所以 k 的取值范围最大是 255，最小是 -255，而我们只需要在这个范围内寻找即可
    for i := -255; i <= 255; i++ {
        // 计算新图像素平均值
        sum := 0
        for _, val := range img {
            newVal := val + i
            if newVal < 0 {
                newVal = 0
            } else if newVal > 255 {
                newVal = 255
            }
            sum += newVal
        }
        newAvg := sum / n
        // 更新最接近128的k
        if abs(newAvg-128) < diff {
            fmt.Println(newAvg)
            diff = abs(newAvg - 128)
            k = i
            fmt.Println(diff, " k:", k)
        }
    }

    // 输出结果
    fmt.Println(k)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
