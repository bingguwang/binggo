package leetCode

import (
	"testing"
)

/**
05-30

反转链表
*/
func TestRsjm(t *testing.T) {
	array := MakeLinkListByArray([]int{1, 2, 3, 4})
	list := reverseList(array)
	PrintLinkList(list)
}

// 三个指针，一个保证未遍历到的不会丢失，一个维护翻转后的链表，一个用于遍历
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	p := head
	for p != nil { // p==nil时，pre指的就是新的头
		node := p.Next // 保存cur.next防止丢失
		p.Next = pre
		pre = p  // pre往后移动移动
		p = node // p往后移动移动
	}
	return pre
}
