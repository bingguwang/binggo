package main

/*
*

	两个栈实现一个队列

队列是一个先进先出的
栈是先进后出的，两个栈分别处理入队和出队就可以实现一个队列
入队的时候先入栈，然后出队的时候压入到另外一个栈里，那另外一个栈的栈顶便是最先入队的元素
*/
func main() {
	stack1 = make([]int, 0)
	stack2 = make([]int, 0)
}

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		// 栈2为空时，栈1全部出栈并且压入栈2
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = stack1[0:0]
	}
	// 栈2顶出栈
	val := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return val
}
