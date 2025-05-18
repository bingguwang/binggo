package tree

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestNameccc(t *testing.T) {

}

// 递归的方式求的树高度
func maxDepth(root *TreeNode) int {
	// write code here
	var ans int
	ans = deep(root)
	return ans
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(deep(root.Left), deep(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 非递归，dfs求
// 其实是使用栈来模拟递归的过程
// 逐层遍历树
// 其实这个过程是非递归先序遍历的一个变种
func maxDepthDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	depths := []int{1} // 记录每个节点的深度, 因为不能修改树结构体里的成员，所以这里使用一个数组外维护，更新频率和节点保持一致
	maxHeight := 0

	for len(stack) > 0 {
		// 栈顶节点元素出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 栈顶节点对应的高度出栈
		depth := depths[len(depths)-1]
		depths = depths[:len(depths)-1]

		// 更新最大高度
		if node != nil {
			maxHeight = max(maxHeight, depth)

			// 将左右子节点压入栈，并记录它们的深度
			if node.Left != nil {
				stack = append(stack, node.Left)
				depths = append(depths, depth+1)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
				depths = append(depths, depth+1)
			}
		}
	}
	return maxHeight
}

// 广度优先遍历 BFS， bfs往往是和队列相关
func maxDepthBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 使用队列进行广度优先搜索
	queue := []*TreeNode{root}
	height := 0

	for len(queue) > 0 {
		// 处理当前层的所有节点
		levelSize := len(queue) // 每层有多少节点事先记录下来
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // 队首是父节点，队首出队时把所有的子节点都加入队列

			// 将当前节点的子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 每处理完一层，高度加 1
		height++
	}
	return height
}
