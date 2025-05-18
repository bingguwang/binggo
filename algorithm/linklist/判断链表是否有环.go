package main

import "fmt"

/**

判断是否有环
通过快慢指针来实现
a+n*(b+c)+b = 2*(a+b)
a=(n-1)*(b+c) + c

由此可以看到，
其实 b+c 正好等于环的长度，也就是说：
从链表头部到入环的距离（a）恰好等于从相遇点到入环点的距离（c）再加上 n-1 圈个环的长度

我们再定义一个指针指向链表起点，一次走一步，slow 指针也同步继续往后走，那么这两个指针就一定会在链表的入口位置相遇。
*/

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	// write code here
	fast, slow := pHead, pHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 到达相遇点了
			// 新的指针从头开始一步步走
			newp := pHead
			for slow != newp {
				newp = newp.Next
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}
func main() {
	solution := MoreThanHalfNum_Solution([]int{1, 2, 3, 2, 2, 2, 5, 4, 2})
	fmt.Println(solution)
}

func MoreThanHalfNum_Solution(numbers []int) int {
	target := numbers[0]
	times := 0
	for i := 0; i < len(numbers); i++ {
		if times == 0 {
			target = numbers[i]
			times++
		} else {
			if target == numbers[i] {
				times++
			} else {
				times--
			}
		}
	}
	fmt.Println("times---	", times)
	return target
}
