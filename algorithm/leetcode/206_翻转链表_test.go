package leetCode

import (
	"testing"
)

/*
*
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
	cur := head
	for cur != nil { // p==nil时，pre指的就是新的头
		next := cur.Next // 保存cur.next防止丢失
		cur.Next = pre
		pre = cur  // pre往后移动移动
		cur = next // cur往后移动移动
	}
	return pre
}
