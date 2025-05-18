package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// str := "128#0#255#255"
	// str := "100#101#1#5"
	str := "1#0#0#0"
	s := strings.Split(str, "#")
	var hexres string
	for _, v := range s {
		// 进行一些校验工作

		// 转为10进制
		dec, _ := strconv.ParseInt(v, 10, 64)
		// 10进制转为16进制字符串
		tmp := strconv.FormatInt(dec, 16)
		//拼接
		if len(tmp) == 1 {
			tmp = "0" + tmp
		}
		hexres = hexres + tmp
		fmt.Println(hexres)
	}
	// 转为10进制
	res, _ := strconv.ParseInt(hexres, 16, 64)
	fmt.Println(res)
}
