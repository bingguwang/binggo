package Backtracking

import (
    "fmt"
)

/**
4 5
X X X X X
O O X O X
X O X O X
X X X O X
*/
var (
    m, n    int
    mapArr  [][]string
    Pos     [2]int // 每个位置
    count   int
    visited [][]bool
)

func main() {
    fmt.Scan(&m, &n)
    mapArr = make([][]string, m)
    for i := 0; i < m; i++ {
        mapArr[i] = make([]string, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&mapArr[i][j])
        }
    }
    visited = make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }
    maxSize := 0     // 找到的最大区域大小
    var quyu [][]int // 每个元素 统计最大区域的入口坐标x y以及其区域大小
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mapArr[i][j] == "O" && !visited[i][j] { // 入口
                visited[i][j] = true         // 替换坐标
                position := make([][]int, 0) // 位置集合
                f(i, j, &position)           // position存储的是当前区域内的所有点的坐标
                if count == 1 {              // 代表当前区域只有一个入口， count记录的当前区域的入口数
                    if maxSize == len(position) { // 和当前最大一样大，最大面积对应了多个区域，则我们不需要记录位置只需要知道最大值
                        quyu = nil
                    } else if maxSize < len(position) { // 比当前最大面积要大，则记录替换
                        quyu = [][]int{{Pos[0], Pos[1], len(position)}}
                        maxSize = len(position)
                    }
                    fmt.Println("一个解：", Pos[0], Pos[1], len(position))
                }
            }
            // 重置参数
            count = 0
            Pos = [2]int{0, 0}
        }
    }
    fmt.Println(maxSize)
    if len(quyu) == 1 { // 说明只有一个最大面积的唯一区域
        outRes := quyu[0]
        fmt.Printf("%d %d %d", outRes[0], outRes[1], outRes[2])
    } else if maxSize != 0 {
        fmt.Println(maxSize)
    } else {
        fmt.Println("NULL")
    }
}

func f(x, y int, list *[][]int) {
    // 判断是否为入口
    if x == 0 || y == 0 || x == m-1 || y == n-1 {
        count++    // 在一个完整的递归链里，应该只有一次节点会在边界上，就是入口的那次，又多次就说明不是空闲区
        Pos[0] = x // 是入口，把坐标保存到Pos
        Pos[1] = y
        *list = append(*list, []int{x, y})
    }
    if x < m-1 {
        if mapArr[x+1][y] == "O" && !visited[x+1][y] {
            visited[x+1][y] = true
            *list = append(*list, []int{x + 1, y})
            f(x+1, y, list)
        }
    }
    if y < n-1 {
        if mapArr[x][y+1] == "O" && !visited[x][y+1] {
            visited[x][y+1] = true
            *list = append(*list, []int{x, y + 1})
            f(x, y+1, list)
        }
    }
    if x > 0 {
        if mapArr[x-1][y] == "O" && !visited[x-1][y] {
            visited[x-1][y] = true
            *list = append(*list, []int{x - 1, y})
            f(x-1, y, list)
        }
    }
    if y > 0 {
        if mapArr[x][y-1] == "O" && !visited[x][y-1] {
            visited[x][y-1] = true
            *list = append(*list, []int{x, y - 1})
            f(x, y-1, list)
        }
    }
}
