package od

import (
    "fmt"
    "sort"
    "testing"
)

/**
给定一个N行M列的二维矩阵，矩阵中每个位置的数字取值为0或1。矩阵示例如：
1 1 0 0
0 0 0 1
0 0 1 1
1 1 1 1

现需要将矩阵中所有的1进行反转为0，规则如下：
1）当点击一个1时，该1变被反转为0，同时相邻的上、下、左、右，以及左上、左下、右上、右下8个方向的1（如果存在1）均会自动反转为0；
2）进一步地，一个位置上的1被反转为0时，与其相邻的8个方向的1（如果存在1）均会自动反转为0；
按照上述规则示例中的矩阵只最少需要点击2次后，所有值均为0。请问，给定一个矩阵，最少需要点击几次后，所有数字均为0？
*/

// 考察dfs 回溯算法

func main() {
    arr := [][]int{
        {1, 1, 0, 0},
        {0, 0, 0, 1},
        {0, 0, 1, 1},
        {1, 1, 1, 1},
    }
    click := minClick(arr, 4, 4)
    fmt.Println(click)
}
func minClick(arr [][]int, n, m int) int {
    var dfs func(i, j int)
    dfs = func(i, j int) {                // i,j是当前中心位置的坐标
        if i-1 >= 0 && arr[i-1][j] == 1 { // 左边元素是1，则左边元素需要进行消除操作
            arr[i-1][j] = 0
            dfs(i-1, j) // 以i-1,j为中心，继续消除操作
        }
        if i-1 >= 0 && j-1 >= 0 && arr[i-1][j-1] == 1 {
            arr[i-1][j-1] = 0
            dfs(i-1, j-1)
        }
        if j+1 < m && i-1 >= 0 && arr[i-1][j+1] == 1 {
            arr[i-1][j+1] = 0
            dfs(i-1, j+1)
        }
        if i+1 < n && arr[i+1][j] == 1 {
            arr[i+1][j] = 0
            dfs(i+1, j)
        }
        if i+1 < n && j+1 < m && arr[i+1][j+1] == 1 {
            arr[i+1][j+1] = 0
            dfs(i+1, j+1)
        }
        if i+1 < n && j-1 >= 0 && arr[i+1][j-1] == 1 {
            arr[i+1][j-1] = 0
            dfs(i+1, j-1)
        }
        if j-1 >= 0 && arr[i][j-1] == 1 {
            arr[i][j-1] = 0
            dfs(i, j-1)
        }
        if j+1 < m && arr[i][j+1] == 1 {
            arr[i][j+1] = 0
            dfs(i, j+1)
        }
    }

    var minclick int
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            // 每个1点都作为消除中心,我们消除手动点击的消除操作只能在1点
            if arr[i][j] == 1 {
                arr[i][j] = 0
                dfs(i, j) // 一次消除操作
                minclick++
            }
        }
    }
    return minclick
}
func TestKsadsad(t *testing.T) {
    a:=[]string{"ADS","aDs"}
    sort.Strings(a)
    fmt.Println(a)

}