package main

import "fmt"

//	func main() {
//		pHead1 := &ListNode{Val: 1}
//		b := &ListNode{Val: 3}
//		c := &ListNode{Val: 5}
//		pHead1.Next = b
//		b.Next = c
//
//		pHead2 := &ListNode{Val: 2}
//		f := &ListNode{Val: 4}
//		g := &ListNode{Val: 6}
//		pHead2.Next = f
//		f.Next = g
//		Merge(pHead1, pHead2)
//	}
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	var head *ListNode
	if pHead1.Val < pHead2.Val {
		head = pHead1
	} else {
		head = pHead2
	}
	p := head
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			p.Next = pHead1
			p = pHead1
			pHead1 = pHead1.Next
		} else {
			p.Next = pHead2
			p = pHead2
			pHead2 = pHead2.Next
		}
	}
	printlist(head)

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printlist(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
}
