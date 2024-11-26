package leetCode

import (
    "fmt"
    "testing"
)

/**

删除链表倒数第n个节点，返回链表
*/
func TestKjs(t *testing.T) {
    list := MakeLinkListByArray([]int{1, 2, 3, 4, 5})

    end := removeNthFromEnds(list, 5)
    PrintLinkList(end)
}

// 遍历两次，先获得链表的长度
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length, i, p := 0, 0, head
    for p != nil {
        length++
        p = p.Next
    }
    fmt.Println(length)
    p = head
    if length == n { // 要删除的是head
        return p.Next
    }
    for {
        i++
        if i == length-n {
            break
        } else {
            p = p.Next
        }
    }
    fmt.Println(p.Val)
    p.Next = p.Next.Next
    return head
}

// 双指针,两个距离为n的指针，这样当前面的指针到末尾的时候，刚好后面的指针就是要删的
func removeNthFromEnds(head *ListNode, n int) *ListNode {
    slow, i, fast := head, 0, head
    for {
        if i == n {
            break
        } else {
            fast = fast.Next
        }
        i++
    }
    if fast == nil { // 要删除的是head
        return head.Next
    }
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    fmt.Println(slow.Val)
    slow.Next = slow.Next.Next
    return head
}

/**
n=2
1 2 3 4 5
s   f
*/

func Del(head *ListNode, n int) *ListNode {
    var size int
    g := head
    for g != nil {
        g = g.Next
        size++
    }
    pre := size - n // // p是要删除的节点的前一个节点
    if pre == 0 {
        return head.Next
    }
    p := head // p是要删除的节点的前一个节点
    for i := 1; i < pre; {
        p = p.Next
        i++
    }
    p.Next = p.Next.Next
    return head
}

func TestKjsss(t *testing.T) {
    list := MakeLinkListByArray([]int{1, 2, 3, 4, 5})

    end := Del(list, 2)
    PrintLinkList(end)
}
