package od

import (
    "fmt"
    "testing"
)

/**
微服务的集成测试

给你一个nxn 的二维矩阵 useTime，其中 useTime[i][i]=10 表示服务i自身启动加载需要消耗10s，
useTime[i][j]=1 表示服务i 启动依赖服务i 启动完成，useTime[i][k]=0，
表示服务i 启动不依赖服务 k其实 0<= i,j，k< n。
服务之间启动没有循环依赖(不会出现环)，若想对任意一个服务i进行集成测试(服务追身也需要加载)，求最少需要等待多少时间。

3
5 0 0
1 5 0
0 1 5
3
*/
var (
    n       int
    visited []bool
    useTime [][]int
    result  []int
    res     int
)

func TestNamse(t *testing.T) {
    fmt.Scan(&n)
    useTime = make([][]int, n)
    for i := 0; i < n; i++ {
        useTime[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&useTime[i][j])
        }
    }
    var k int
    fmt.Scan(&k)
    visited = make([]bool, n)

    i := solution(useTime, k-1)
    fmt.Println(i)

}
func solution(array [][]int, row int) int {
    rowArray := array[row]
    // 当前要计算节点
    sum := array[row][row] // sum保存此节点需要的时间
    max := 0
    // 遍历所有此节点依赖的节点，计算row依赖的每个节点的所需用时，找出最大用时max
    for colum := 0; colum < len(rowArray); colum++ {
        if row != colum && array[row][colum] == 1 { // 当前节点依赖column
            // 递归计算colum需要的用时value
            value := solution(array, colum)
            if value > max {
                max = value
            }
        }
    }
    sum += max // sum就是row需要的最大用时了
    for colum := 0; colum < len(rowArray); colum++ {
        if row != colum {
            array[row][colum] = 0
        }
    }
    // 更新一下节点需要的用时
    array[row][row] = sum
    fmt.Println(array)
    return sum
}
