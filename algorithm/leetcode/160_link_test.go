package leetCode

import (
    "fmt"
    "testing"
)

/**
05-09 上午
*/
/**
两个链表的相交点
*/
func TestSa(t *testing.T) {
    a1 := &ListNode1{Val: 1}
    a2 := &ListNode1{Val: 2}
    a3 := &ListNode1{Val: 3}
    a1.Next = a2
    a2.Next = a3

    b1 := &ListNode1{Val: 3}
    b2 := &ListNode1{Val: 4}
    b1.Next = b2
    getIntersectionNode(a1, b1)

}

func getIntersectionNode(headA, headB *ListNode1) *ListNode1 {
    if headA == nil || headB == nil {
        return nil
    }
    a, b := headA, headB
    for a != b { // 未相交
        if a == nil { // a链表遍历完，就从b链表头开始继续遍历
            a = headB
        } else {
            a = a.Next
        }
        if b == nil { // b链表遍历完，就从a链表头开始继续遍历
            b = headA
        } else {
            b = b.Next
        }
    }
    fmt.Println(a == nil)
    fmt.Println(b == nil)
    return a
}
type ListNode1 struct {
    Val  int
    Next *ListNode1
}