package leetCode

import (
    "testing"
)

func TestSum(t *testing.T) {
    l1 :=  MakeLinkListByArray([]int{9,9,9})
    l2 :=  MakeLinkListByArray([]int{9,9})
     PrintLinkList(l1)
     PrintLinkList(l2)
    l := addTwoNumbers(l1, l2)
     PrintLinkList(l)
}
func addTwoNumbers(l1 * ListNode, l2 * ListNode) * ListNode {
    head := & ListNode{}
    p := head
    carry := 0
    for l1 != nil || l2 != nil {
        a, b := 0, 0
        if l1 == nil {
            a = 0
        } else {
            p.Next = l1
            a = l1.Val
            l1 = l1.Next
        }
        if l2 == nil {
            b = 0
        } else {
            b = l2.Val
            l2 = l2.Next
        }
        p.Next = & ListNode{Val: (a + b + carry) % 10}
        carry = (a + b + carry) / 10
        p = p.Next
    }
    if carry >0 {
        p.Next = & ListNode{Val: carry}
    }

    return head.Next
}
