package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//inputArray()
	inputArr()
}

// 输入数组长宽，然后输入数组元素
func inputArray() {
	scanner := bufio.NewScanner(os.Stdin)

	// 输入矩阵的长和宽
	fmt.Print("请输入矩阵的长和宽（用逗号分隔）: ")
	scanner.Scan()
	dims := strings.Split(scanner.Text(), ",")
	if len(dims) != 2 {
		fmt.Println("输入格式错误，请使用 '长,宽' 的格式")
		return
	}

	rows, err1 := strconv.Atoi(dims[0])
	cols, err2 := strconv.Atoi(dims[1])
	if err1 != nil || err2 != nil {
		fmt.Println("输入的长和宽必须是整数")
		return
	}

	// 创建矩阵
	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
	}

	// 输入矩阵内容
	for i := 0; i < rows; i++ {
		fmt.Printf("请输入第 %d 行的内容（用空格分隔）: ", i+1)
		scanner.Scan()
		//elements := strings.Fields(scanner.Text())
		elements := strings.Split(scanner.Text(), ",")
		if len(elements) != cols {
			fmt.Printf("第 %d 行的内容长度不正确，应为 %d 个元素\n", i+1, cols)
			return
		}
		for j := 0; j < cols; j++ {
			matrix[i][j] = elements[j]
		}
	}

	// 输出矩阵
	fmt.Println("输入的矩阵是:")
	for _, row := range matrix {
		fmt.Println(strings.Join(row, " "))
	}
}

// 上来直接输入数组,这种一般是事先知道数组的长宽的，所以可以控制循环何时退出
func inputArr() {

	scanner := bufio.NewScanner(os.Stdin)
	var i = 0
	count := 0 // 用于控制下面循环的退出
	for scanner.Scan() {
		text := scanner.Text()
		i++
		fields := strings.Fields(text)
		fmt.Println(len(fields))
		count = len(fields)
		if i == count {
			break
		}
	}

}
