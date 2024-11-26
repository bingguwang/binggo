package leetCode

import (
    "testing"
)

/**
05-09 上午
*/
/**
判断链表是否有环，不能使用额外的空间。如果有环，输出环的起点指针，如果没有环，则输出空。
*/
func TestGS(t *testing.T) {

}
func detectCycle(head *ListNode) *ListNode {
   slow, fast := head, head
   var meet *ListNode // 相遇点
   for fast != nil && fast.Next != nil {
       slow = slow.Next
       fast = fast.Next.Next
       if fast == slow {
           meet = fast
           break
       }
   }
   if meet == nil { // 没相遇，说明无环
       return nil
   }
   slow = head
   for meet != slow { // 相遇后，各自从的起点，相遇点都用1步长移动，一定会在环起点再次相遇
       meet = meet.Next
       slow = slow.Next
   }
   return meet
}
