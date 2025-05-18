package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
*
给定长度为n的无序的数字数组Q，每个数字代表二叉树的叶子节点的权值，数字数组的值均大于等于 1
请完成一个函数只，根据输入的数字数组，生成哈夫曼树只，并将哈夫曼树按照中序遍历输出。为了保证输出的 二叉树9 中序遍历只结果统一，增加以下限制:
在树节点中，左节点权值小于等于右节点权值，根节点权值为左右节点权值之和。
当左右节点权值相同时，左子树高度高度小于等于右子树。
注意:所有用例保证有效，并能生成哈夫曼树提醒:哈夫曼树又称最优二叉树，是一种带权路径长度最短的一叉树。
所谓树的带权路径长度，就是树中所有的叶结点的权值乘上其到根结点的路径长度(若根结点为0层，叶结点到根结点的路径长度为叶结点的层数)。

详解https://blog.csdn.net/weixin_48157259/article/details/142605068

先了解一下什么是哈夫曼树


*/

// 定义树节点的结构体
type TreeNode struct {
	val    int
	left   *TreeNode
	right  *TreeNode
	height int
}

// 构建哈夫曼树的函数
func buildTree(nodeList []*TreeNode) *TreeNode {
	for len(nodeList) > 1 {
		// 对nodeList进行排序，先按照节点值val升序排序，再按照节点高度height升序排序
		sort.Slice(nodeList, func(i, j int) bool {
			if nodeList[i].val == nodeList[j].val {
				return nodeList[i].height < nodeList[j].height
			}
			return nodeList[i].val < nodeList[j].val
		})
		// 弹出前两个元素，为当前所选择的，用来构建新节点的两个节点
		left := nodeList[0]
		right := nodeList[1]
		nodeList = nodeList[2:]
		// 新节点的值为左右节点值的和
		val := left.val + right.val
		// 新节点的高度为左右节点的高度的较大值再+1
		height := max(left.height, right.height) + 1
		// 构建新节点newNode
		newNode := &TreeNode{val: val, left: left, right: right, height: height}
		// 将新节点加入列表nodeList中
		nodeList = append(nodeList, newNode)
	}
	// 返回根节点
	if len(nodeList) == 0 {
		return nil
	}
	return nodeList[0]
}

// 中序遍历的函数
func inorder(ans *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	inorder(ans, node.left)
	*ans = append(*ans, node.val)
	inorder(ans, node.right)
}

// 辅助函数：返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 主函数
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	tmp := strings.Fields(scanner.Text())

	// 初始化节点列表nodeList，包含所有叶节点
	nodeList := make([]*TreeNode, n)

	for i, v := range tmp {
		val, _ := strconv.Atoi(v)
		nodeList[i] = &TreeNode{val: val, left: nil, right: nil, height: 0}
	}

	// 构建树，退出后nodeList的长度为1
	root := buildTree(nodeList)

	// 初始化答案列表
	ans := []int{}
	// 中序遍历函数，传入根节点
	inorder(&ans, root)

	// 输出结果
	for i, val := range ans {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
}
