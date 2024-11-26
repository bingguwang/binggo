package leetCode

import (
    "fmt"
    "testing"
)

/**

二叉树前序遍历
*/
func preorderTraversal(root *TreeNode) []int {
    var res []int
    predOrder(root, res)
    return res
}

func predOrder(root *TreeNode, res []int) {
    if root == nil {
        return
    }
    fmt.Printf("%p\n", res)
    b := append(res, root.Val) // 如果不是传指针，由于是赋值给原切片，所以会新开辟一个底层数组，所以影响不了原来的底层数组
    predOrder(root.Left, b)
    predOrder(root.Right, b)
}

func TestRSD(t *testing.T) {
    a1 := &TreeNode{Val: 1}
    a2 := &TreeNode{Val: 2}
    a3 := &TreeNode{Val: 3}
    a4 := &TreeNode{Val: 4}
    a5 := &TreeNode{Val: 5}
    a1.Left = a2
    a1.Right = a3
    a3.Left = a4
    a3.Right = a5

    //res := Okdsdk(a1)
    res := preorderTraversal(a1)
    fmt.Println(res)

}

func TestJs(t *testing.T) {
    a := []int{1, 2}
    fmt.Printf("%p\n", a)
    func(i []int) { // 传入的是拷贝的切片xx，切片里的data指针也是拷贝的，指针还是指向原来的底层数组
        fmt.Printf("%p\n", i)
        i[0] = 5              // 因为data指针指向的是原来的底层数组，所以这改变会影响到原切片的值
        fmt.Printf("%p\n", i) // 所以这里指针还是原来的data指针
    }(a)
    fmt.Printf("%p\n", a)
    fmt.Println(a) // 所以这里不会受影响还是原来的
}
