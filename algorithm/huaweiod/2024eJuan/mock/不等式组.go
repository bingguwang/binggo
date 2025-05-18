package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	strs := strings.Split(s.Text(), ";")
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
	fmt.Println("-------------")
	var xishuarr [][]float64
	for i := 0; i < len(strs)-3; i++ {
		tmp := make([]float64, 0)
		for _, v := range strings.Split(strings.TrimSpace(strs[i]), ",") {
			t, _ := strconv.ParseFloat(v, 64)
			tmp = append(tmp, t)
		}
		xishuarr = append(xishuarr, tmp)
	}
	fmt.Println(xishuarr)

	var bianliangarr []int
	for _, v := range strings.Split(strings.TrimSpace(strs[len(strs)-3]), ",") {
		t, _ := strconv.Atoi(v)
		bianliangarr = append(bianliangarr, t)
	}

	var tararr []float64
	for _, v := range strings.Split(strs[len(strs)-2], ",") {
		t, _ := strconv.ParseFloat(v, 64)
		tararr = append(tararr, t)
	}
	var fuhaoarr []string
	for _, v := range strings.Split(strs[len(strs)-1], ",") {
		fuhaoarr = append(fuhaoarr, v)
	}
	fmt.Println("变量数组:")
	fmt.Println(bianliangarr)
	fmt.Println("目标值数组:")
	fmt.Println(tararr)
	fmt.Println("符号数组:")
	fmt.Println(fuhaoarr)

	var panduanres = true
	var maxcha = -1000000.0
	for i := 0; i < len(xishuarr); i++ {
		var result float64
		for j := 0; j < len(xishuarr[i]); j++ {
			result += xishuarr[i][j] * float64(bianliangarr[j])
		}
		fmt.Println(result)
		switch {
		case fuhaoarr[i] == "<":
			panduanres = panduanres && result < tararr[i]
		case fuhaoarr[i] == ">":
			panduanres = panduanres && result > tararr[i]
		case fuhaoarr[i] == "<=":
			panduanres = panduanres && result <= tararr[i]
		case fuhaoarr[i] == ">=":
			panduanres = panduanres && result >= tararr[i]
		case fuhaoarr[i] == "==":
			panduanres = panduanres && result == tararr[i]
		}
		if maxcha < result-tararr[i] {
			maxcha = result - tararr[i]
		}
	}
	fmt.Println(panduanres)
	fmt.Println(int(maxcha))
}
