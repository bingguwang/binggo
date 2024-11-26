package Backtracking

import (
    "fmt"
)

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/
func main() {
    res := generateParenthesis(3)
    // ["((()))","(()())","(())()","()(())","()()()"]
    fmt.Println(res)
}
func generateParenthesis(n int) []string {
    var res []string
    var backtrack func(cur string, left, right, max int)
    /**
      res：指向结果切片的指针
      cur：当前正在生成的字符串
      open：当前已经使用的左括号数目
      close：当前已经使用的右括号数目
      max：需要生成的括号对数
    */
    backtrack = func(cur string, left, right, max int) {
        if len(cur) == max*2 { // 有个结果了
            res = append(res, cur)
            return
        }
        if left < max { // 要生成n对括号，左括号数还不够
            backtrack(cur+"(", left+1, right, max)
        }
        if right < left { // 右括号少于左括号， 还需要增加右括号
            backtrack(cur+")", left, right+1, max)
        }
    }
    backtrack("", 0, 0, n)
    return res
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func printListnodes(h *ListNode) {
    for h != nil {
        fmt.Printf("%v--", h.Val)
        h = h.Next
    }
    fmt.Println()
}
