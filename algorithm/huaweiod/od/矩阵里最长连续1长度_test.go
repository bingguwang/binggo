package od

import (
    "fmt"
    "testing"
)
// dfs
func TestFhasd(t *testing.T) {
    var arr [][]int
    arr = [][]int{
        {0, 1, 1, 0},
        {0, 1, 1, 0},
        {1, 1, 0, 1},
    }
    //矩阵的行和列
    len_row, len_col := len(arr), len(arr[0])

    //动态数组，分别保存行，列，对角线和反对角线上的当前元素下连续1的个数
    f := make([][][]int, 4)
    for i := 0; i < 4; i++ {
        rows := make([][]int, len_row+2) // 为什是+2, 看打印结果可以看到，要包住我们的矩阵
        for j := 0; j < len(rows); j++ {
            rows[j] = make([]int, len_col+2)
        }
        f[i] = rows
    }

    for i := 0; i < len(f[0]); i++ {
        for j := 0; j < len(f[0][i]); j++ {
            fmt.Printf("%d ", f[0][i][j])
        }
        fmt.Println()
    }

    res := 0 //保存中间出现过的最长的长度
    //遍历数组元素
    for i := 1; i <= len_row; i++ {
        for j := 1; j <= len_col; j++ {
            if arr[i-1][j-1] == 1 { // 当前字符是1
                f[0][i][j] = f[0][i+0][j-1] + 1 // 行
                f[1][i][j] = f[1][i-1][j+0] + 1 // 列
                f[2][i][j] = f[2][i-1][j-1] + 1 // 对角线
                f[3][i][j] = f[3][i-1][j+1] + 1 // 反对角线

                fmt.Println("=========")
                Prinsad(f[0])

                for k := 0; k < 4; k++ {
                    res = max(res, f[k][i][j])
                }
            } else {
                //更新元素值为0
                for k := 0; k < 4; k++ {
                    f[k][i][j] = 0
                }
            }
        }
    }
    fmt.Println(res)
}

func Prinsad(a [][]int) {
    for j := 0; j < len(a); j++ {
        for i := 0; i < len(a[j]); i++ {
            fmt.Printf("%d ", a[j][i])
        }
        fmt.Println()
    }
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}