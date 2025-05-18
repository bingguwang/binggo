package od

import (
    "fmt"
    "testing"
)

func TestGalwsa(t *testing.T) {

    //borad := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
    borad := [][]byte{
        {'A', 'C', 'C', 'F'},
        {'C', 'D', 'E', 'D'},
        {'B', 'E', 'S', 'S'},
        {'F', 'E', 'C', 'A'},
    }
    //borad := [][]byte{
    //    {'A', 'C'},
    //    {'C', 'D'},
    //}
    b := exist(borad, "ACCESS")
    //b := exist(borad, "ACD")
    fmt.Println(b)
}
func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    used := make([][]bool, m)
    for i := 0; i < m; i++ {
        used[i] = make([]bool, n)
    }
    var path string
    var canFindRamin func(int, int, int) bool
    canFindRamin = func(i, j, p int) bool { // p是word的下标， 指向当前正在查找的字符
        if p == len(word) {
            return true
        }
        // 判断是否越界
        if i < 0 || j < 0 || i >= m || j >= n {
            return false
        }

        if used[i][j] || board[i][j] != word[p] {
            return false
        }
        used[i][j] = true
        path += fmt.Sprintf("%v,%v,", i, j)
        fmt.Println(path)
        // 当前i,j匹配到了当前要找的字符，然后就是要看后面的子递归能否匹配到下一个要找的字符
        canFind := canFindRamin(i+1, j, p+1) || canFindRamin(i-1, j, p+1) || canFindRamin(i, j+1, p+1) || canFindRamin(i, j-1, p+1)
        if canFind { // 子递归可以找到下一个字符
            return true
        } else { // 子递归不能找到下一个要找的字符,则回溯
            used[i][j] = false
            path = path[:len(path)-4] // 回溯的时候path也要回溯
            return false
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] && canFindRamin(i, j, 0) {
                return true
            }
        }
    }
    fmt.Println(path)
    return false
}
