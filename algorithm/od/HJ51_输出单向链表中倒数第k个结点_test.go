package od

import (
    "fmt"
    "testing"
)

type ListNode1 struct {
    m_nKey  int
    m_pNext *ListNode1
}

func TestG(t *testing.T) {
    for {
        var n int
        _, err := fmt.Scan(&n)
        if err != nil { // EOF时会退出循环
            break
        }
        nodes := make([]*ListNode1, n)
        for i := 0; i < n; i++ {
            nodes[i] = &ListNode1{}
            fmt.Scan(&nodes[i].m_nKey)
        }
        for i := 0; i < n-1; i++ {
            nodes[i].m_pNext = nodes[i+1]
        }
        var k int
        fmt.Scan(&k)
        // 双指针
        a := nodes[0]
        b := nodes[0]
        var i int
        for {
            i++
            b = b.m_pNext
            if k == i { // 指针b移到第K个
                break
            }
        }

        for b != nil { // b到尾时，a所在的就是所求
            a = a.m_pNext
            b = b.m_pNext
        }
        fmt.Println(a.m_nKey)
    }

}
