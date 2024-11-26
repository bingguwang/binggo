package leetCode

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func MakeLinkListByArray(a []int) *ListNode {
    head := &ListNode{}
    p := head
    for _, v := range a {
        p.Next = &ListNode{Val: v}
        p = p.Next
    }
    return head.Next
}

func PrintLinkList(h *ListNode) {
    for h != nil {
        fmt.Printf("->%v", h.Val)
        h = h.Next
    }
    fmt.Println()
}
