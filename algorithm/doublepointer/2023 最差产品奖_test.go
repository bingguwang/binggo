package doublepointer

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

/**
A公司准备对他下面的N个产品评选最差奖，评选的方式是首先对每个产品进行评分，然后根据评分区间计算相邻几个产品中最差的产品。
评选的标准是依次找到从当前产品开始前M个产品中最差的产品，请给出最差产品的评分序列。

第一行，数字M，表示评分区间的长度，取值范围是0<M<10000
第二行，产品的评分序列，比如[12,3,8,6,5]，产品数量N范围是-10000<N<10000
输入：
3
12,3,8,6,5
输出：
3,3,5
说明：
12,3,8 最差的是3
3,8,6 中最差的是3
8,6,5 中最差的是5

// 思路：双指针
确定好滑动窗口，在窗口内找出最差值

3
12,3,8,6,5
4
12,3,8,6,5,12,3,8,6,5
*/
func main() {
    var m int
    fmt.Scan(&m)

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    split := strings.Split(text, ",")

    a := make([]int, len(split))
    for i := 0; i < len(split); i++ {
        atoi, _ := strconv.Atoi(split[i])
        a[i] = atoi
    }

    var res []int
    for i := 0; i <= len(a)-m; i++ {
        min := math.MaxInt
        for j := i; j <= i+m-1; j++ {
            if min > a[j] {
                min = a[j]
            }
        }
        res = append(res, min)
    }
    for i := 0; i < len(res); i++ {
        fmt.Printf("%v", res[i])
        if i != len(res)-1 {
            fmt.Printf(",")
        }
    }

}
