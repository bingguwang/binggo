package leetCode

import (
    "fmt"
    "testing"
)

/**

反转链表，按指定区间反转
*/
func TestFa(t *testing.T) {
    list := MakeLinkListByArray([]int{1, 2, 3, 4, 5})
    node := reverseBetween(list, 3, 5)
    PrintLinkList(node)
}

// 头插法来实现反转，3指针，一个cur(一直指向反转区的第一个节点)(一定就跟着移动)，一个pre一直指向反转区的前一个非反转区节点，一个next指向a的下一个节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummyNode := &ListNode{Val: -1, Next: head}
    pre := dummyNode // 设一个虚节点是因为当反转区是head开始的时候，pre就需要一个空节点了

    for i := 0; i < left-1; i++ {
        pre = pre.Next
    }
    fmt.Println(pre.Val)
    cur := pre.Next
    i := 0
    for i != right-left { // 每次的next指向的节点就是本次翻转后要扭转到的反转区头部的节点
        next := cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
        fmt.Println(next.Val)
        PrintLinkList(head)
        i++
    }

    return dummyNode.Next
}

func ewqreverseBetween(head *ListNode, left int, right int) *ListNode {
    dumHead := &ListNode{Next: head}
    pre := dumHead
    for i := 1; i < left; i++ {
        pre = pre.Next
    }
    fmt.Println(pre.Val)
    return nil
}

func TestGja(t *testing.T) {
    list := MakeLinkListByArray([]int{1, 2, 3, 4, 5})
    node := ewqreverseBetween(list, 3, 4)
    PrintLinkList(node)
}
