package leetCode

import (
    "fmt"
    "testing"
)

/**
05-09 上午
*/
// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

func TestSas(t *testing.T) {

}

func inorderTraversal(root *TreeNode) []int {
    var res []int
    inorder(root, &res)
    inorderTraversal(root)
    return res
}
func inorder(root *TreeNode, arr *[]int) {
    if root != nil {
        inorder(root.Left, arr)
        *arr = append(*arr, root.Val)
        inorder(root.Right, arr)
    }
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func TestUSD(t *testing.T)  {
    a1:=&TreeNode{Val: 1}
    a2:=&TreeNode{Val: 2}
    a3:=&TreeNode{Val: 3}
    a4:=&TreeNode{Val: 4}
    a5:=&TreeNode{Val: 5}
    a1.Left = a2
    a1.Right = a3
    a3.Left = a4
    a3.Right = a5

    //res := Okdsdk(a1)
    res := Okdsdk(nil)
    fmt.Println(res)
}
func Okdsdk(node *TreeNode) []int {
    var res []int
    Ssod(&res, node)
    fmt.Println(res)
    return res
}

func Ssod(res *[]int, node *TreeNode) {
    if node == nil {
        return
    }
    Ssod(res, node.Left)
    fmt.Println(node.Val)
    *res = append(*res, node.Val)
    Ssod(res, node.Right)
}
