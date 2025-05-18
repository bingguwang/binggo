package main

import (
	"bufio"
	"fmt"
	"github.com/zeromicro/go-zero/core/mathx"
	"os"
	"strconv"
	"strings"
)

/*
*
题目描述
给定一个二叉树，每个节点上站着一个人，节点数字表示父节点到该节点传递悄悄话需要花费的时间。
初始时，根节点所在位置的人有一个悄悄话想要传递给其他人，求二叉树所有节点上的人都接收到悄悄话花费的时间。
输入描述
给定一个数组表示二叉树，-1 表示空节点
输出描述
返回所有节点都接收到悄悄话花费的时间
输入
0 9 20 -1 -1 15 7 -1 -1 -1 -1 3 2
输出
38
*/
type treenode struct {
	val   int
	left  *treenode
	right *treenode
}

// 首先还是要想办法还原这棵树
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	list := make([]int, 0)
	for _, v := range fields {
		val, _ := strconv.Atoi(v)
		list = append(list, val)
	}
	var nodelist = make([]*treenode, len(list))
	for i, v := range list {
		if v != -1 {
			nodelist[i] = &treenode{val: v}
		}
	}
	for i, v := range list {
		if v != -1 {
			if i*2+1 < len(list) { //左子树
				nodelist[i].left = nodelist[i*2+1]
			}
			if i*2+2 < len(list) { //右子树
				nodelist[i].right = nodelist[i*2+2]
			}
		}
	}
	// 根节点
	root := nodelist[0]
	// 之后的问题感觉就不是简单的dfs，要求的结果好像还有一点动态规划的赶脚了
	dfss(0, root)
	fmt.Println(anss)
	//fmt.Println(dfsXtoS(root))
}

var anss int

// 如果不考虑动态规划的话，就从上而下，暴力尝试每一条路径，最后最大时间的那条路径就是要找的路径
func dfss(presum int, nod *treenode) {
	cur := presum + nod.val
	// 遇到叶子节点，更新结果
	if nod.right == nil && nod.left == nil {
		anss = mathx.MaxInt(cur, anss)
		return
	}
	if nod.left != nil { // 递归左子树
		dfss(cur, nod.left)
	}
	if nod.right != nil { // 递归右子树
		dfss(cur, nod.right)
	}
}

// 自下而上就有点动态规划的赶脚
// 自下而上还是自下而上，其实看终止条件就可以大概区分了
func dfsXtoS(nod *treenode) int {
	// 遇到空节点返回0,表示此时节点对结果没有贡献
	if nod == nil {
		return 0
	}
	leftsum := dfsXtoS(nod.left)
	rightsum := dfsXtoS(nod.right)
	return mathx.MaxInt(leftsum, rightsum) + nod.val
}
