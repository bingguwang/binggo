package od

import (
    "fmt"
    "sort"
    "testing"
)

// 考察 栈 回溯算法 dfs

func TestTrain(t *testing.T) {
    var lth int
    fmt.Scan(&lth)
    trains := make([]int, lth)

    for i := 0; i < lth; i++ {
        fmt.Scan(&trains[i])
    }
    var stack, res []int
    var ress []string
    var dfs func(int)
    dfs = func(i int) { // i是需要决定是否入栈的火车编号
        if len(res) == lth {
            fmt.Println(res)
            s := ""
            for i := 0; i < len(res); i++ {
                s += fmt.Sprintf("%d ", res[i])
            }
            ress = append(ress, s)
            return
        }

        // 因为进站和出站这2种选择是并列的
        // 所以当选了第一种时，操作完后需要将数据恢复，以保证选2和选1的初始条件是相同的
        if i < lth { // 可以选择一个进站
            // 火车入栈
            fmt.Println(trains[i], "入栈")
            stack = append(stack, trains[i])
            dfs(i + 1)                   // 子递归去让下一个火车入栈
            stack = stack[:len(stack)-1] // 回溯，撤销掉“ 选择让火车入栈”导致的结果
            fmt.Println("恢复栈,此时栈内:", stack)
        }
        if len(stack) != 0 { // 选择栈内火车出栈
            fmt.Println(stack[len(stack)-1], "出栈")
            // 记录栈顶。栈顶出栈
            res = append(res, stack[len(stack)-1])
            stack = stack[:len(stack)-1]

            dfs(i) // 因为选择的是火车出栈，i没有进入栈内，所以子递归还是i
            // 回溯，撤销"选择栈内火车出栈"造成的结果
            stack = append(stack, res[len(res)-1])
            res = res[:len(res)-1]
        }
    }
    dfs(0)

    sort.Strings(ress)
    for _, v := range ress {
        fmt.Println(v)
    }

}
