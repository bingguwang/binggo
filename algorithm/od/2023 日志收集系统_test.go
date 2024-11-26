package od

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    split := strings.Split(text, " ")
    num := make([]int, len(split))
    for i := 0; i < len(split); i++ {
        num[i], _ = strconv.Atoi(split[i])
    }
    markMax := -1
    count := 0
    //最优情况，第一个数就足够上报
    if num[0] >= 100 {
        fmt.Println(100)
        return
    }

    for i := 0; i < len(num); i++ {
        count += num[i]
        if count >= 100 { // count用于判断是否需要上报了
            markMax = max(markMax, 100-mark(i, num))
            break
        } else {
            markMax = max(markMax, count-mark(i, num))
        }
    }

    fmt.Println(markMax)
}

func mark(n int, num []int) int { // 下标是n以及n之前的数都是延迟了的，返回的是需要扣的分数
    sum := 0
    for i := 0; i <= n; i++ {
        sum += num[i] * (n - i) // n-i是延迟了多少秒，每延迟1秒扣1分
    }
    fmt.Println("sum:", sum)
    return sum
}
