package Backtracking

import (
    "fmt"
    "testing"
)

func TestHkskdlq(t *testing.T) {
    arr := [][]int{
        {0, 1, 0, 0, 0},
        {0, 1, 0, 1, 0},
        {0, 0, 0, 1, 1},
        {0, 1, 0, 0, 0},
        {0, 0, 0, 1, 0},
    }

    m, n := len(arr), len(arr[0])
    used := make([][]bool, m)
    for i := 0; i < m; i++ {
        used[i] = make([]bool, n)
    }
    var path []string
    var canMake func(int, int) bool
    canMake = func(i, j int) bool {
        if i == m-1 && j == n-1 {
            path = append(path, fmt.Sprintf("(%d,%d)", i, j))
            return true
        }
        // 越界?
        if i >= m || j >= n || i < 0 || j < 0 {
            return false
        }
        // 当前位置能走？
        if used[i][j] || arr[i][j] == 1 {
            return false
        }
        used[i][j] = true
        path = append(path, fmt.Sprintf("(%d,%d)", i, j))
        canRes := canMake(i+1, j) || canMake(i-1, j) || canMake(i, j+1) || canMake(i, j-1)
        if canRes {
            return true
        } else {
            // 回溯
            used[i][j] = false
            path = path[:len(path)-1]
            return false
        }
    }

    canMake(0, 0)
    fmt.Println(path)
}