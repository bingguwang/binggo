package leetCode

import (
    "testing"
)

func TestSA(t *testing.T) {

}

/**
05-09 上午
*/
/**
给你一个链表的头节点 head ，判断链表中是否有环。
*/

func hasCycle(head *ListNode) bool {
   // 判断是否含环，龟兔算法，快慢指针
   fast := head
   slow := head
   for fast != nil && fast.Next != nil {
       fast = fast.Next.Next
       slow = slow.Next
       if fast == slow { // 如果二者可以在某处相遇，则一定有环
           return true
       }
   }
   // 快指针如果无法走下去，就说明到达 了的终点，不存在环
   return false
}

