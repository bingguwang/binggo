package od

import (
    "fmt"
)

/**
给定一个由纯数字组成以字符串表示的数值，现要求字符串中的每个数字最多只能出现2次，超过的需要进行删除；删除某个重复的数字后，其它数字相对位置保持不变。
如"34533"，数字3重复超过2次，需要删除其中一个3，删除第一个3后获得最大数值"4533"
请返回经过删除操作后的最大的数值，以字符串表示。

// 思路：
因为删除一个数之后，后面的数需要往前移一位，所以需要关注的是删除掉的这个数，它后面的这个数是不是比它大
如果大的话，后面的数移上来之后，可以保证整个数是变大的，
于是，我们只需要关注两个位置，就是某个过量的数字，出现的第一个位置后最后一个位置
*/
func main() {
    var str string
    fmt.Scan(&str)

    num := make([]int, 10)
    reslist := make([]int, 0)

    for i := 0; i < len(str); i++ {
        x := int(str[i] - '0')
        if num[x] < 2 {
            fmt.Println("加入", x)
            reslist = append(reslist, x)
            num[x]++
        } else {
            val1 := -1
            for j := 0; j < len(reslist); j++ {
                if reslist[j] == x {
                    val1 = j
                    break
                }
            }
            val2 := -1
            for j := len(reslist) - 1; j >= 0; j-- {
                if reslist[j] == x {
                    val2 = j
                    break
                }
            }
            if val1+1 < len(reslist) && reslist[val1+1] > reslist[val1] {
                tmp := reslist[val1+1:]
                reslist = reslist[:val1]
                reslist = append(reslist, tmp...)
                reslist = append(reslist, x)
            } else if val2+1 >= 0 && reslist[val2+1] > reslist[val2] {
                tmp := reslist[val2+1:]
                reslist = reslist[:val2]
                reslist = append(reslist, tmp...)
                reslist = append(reslist, x)
            }
        }
    }

    for i := 0; i < len(reslist); i++ {
        fmt.Print(reslist[i])
    }
}
