package od

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

/**
有一个特异性的双端队列，该队列可以从头部或尾部添加数据，但是只能从头部移出数据。
小A依次执行2n个指令往队列中添加数据和移出数据。其中n个指令是添加数据（可能从头部添加、也可能从尾部添加），依次添加1到n；n个指令是移出数据。现在要求移除数据的顺序为1到n。为了满足最后输出的要求，小A可以在任何时候调整队列中数据的顺序。
请问 小A 最少需要调整几次才能够满足移除数据的顺序正好是1到n；
3
head add 1
remove
tail add 2
head add 3
remove
remove
*/
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var n int
    fmt.Scan(&n)
    deque := make([]int, 0)
    index := 1 // index是要输出的数
    res := 0   // 记录调整的次数
    for i := 0; i < 2*n; i++ {
        scanner.Scan()
        str := strings.Split(scanner.Text(), " ")
        if len(str) == 1 { // remove操作
            if len(deque) == 0 {
                continue
            }
            if deque[0] != index { // 首元素是要输出的元素
                sort.Ints(deque) // 排序之后就是了
                deque = deque[:0]
                res++ // 调整一次
            }
            if len(deque) > 0 {
                deque = deque[1:] // 移除头元素
            } else {
                deque = []int{}
            }
            index++
        } else {
            num, _ := strconv.Atoi(str[2])
            if str[0] == "head" { // 头部插入
                deque = append([]int{num}, deque...)
            } else {
                deque = append(deque, num)
            }
        }
    }
    fmt.Println(res)
}
