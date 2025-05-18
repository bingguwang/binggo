package main

func ReverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next // 保存下一个节点
		cur.Next = prev  // 当前节点反转
		prev = cur
		cur = next // 更新cur
	}
	return prev

}
