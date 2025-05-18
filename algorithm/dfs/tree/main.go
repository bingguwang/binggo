package main

import "fmt"

/*
*
树的遍历

对于的树的遍历，有前中后三序遍历
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 前序
func preOrderDFS(root *TreeNode) {
	if root == nil {
		return
	}

	// 访问当前节点
	processNode(root)

	// 递归遍历左子树
	preOrderDFS(root.Left)

	// 递归遍历右子树
	preOrderDFS(root.Right)
}

func processNode(node *TreeNode) {
	// 在这里处理节点，比如打印节点的值
	fmt.Println(node.Val)
}

// 中序
func inOrderDFS(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归遍历左子树
	inOrderDFS(root.Left)

	// 访问当前节点
	processNode(root)

	// 递归遍历右子树
	inOrderDFS(root.Right)
}

// 后序遍历
func postOrderDFS(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归遍历左子树
	postOrderDFS(root.Left)

	// 递归遍历右子树
	postOrderDFS(root.Right)

	// 访问当前节点
	processNode(root)
}

// 采用dfs建一棵树
func dfsBuildTree() {

}

// 非递归，树的后序遍历
func postOrder(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	visited := make(map[*TreeNode]bool)
	var stack = make([]*TreeNode, 0)
	var res []int
	// 确保左右子树都访问过了，再处理本节点
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		if (root.Right == nil || visited[root.Right]) && (root.Left == nil || visited[root.Left]) {
			// 叶子节点，或者左右子树都访问过了
			// 处理当前节点（后序遍历）
			fmt.Println(root.Val)
			res = append(res, root.Val)
			visited[root] = true
		} else {
			//将左右子节点压入栈，注意顺序
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}
	return res
}
