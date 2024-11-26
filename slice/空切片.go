package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
注意一下切片实际是一个结构：

	type SliceHeader struct {
	 Data uintptr  //引用数组指针地址
	 Len  int     // 切片的目前使用长度
	 Cap  int     // 切片的容量
	}

nil的引用数组的地址是0，也就是没有指向任何实际地址
空切片指向的引用数组地址是一样的，是同一个固定值。所有的空切片指向的数组引用地址都是一样的
*/
func main() {

	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)

	fmt.Printf("s1 pointer:%+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("s2 pointer:%+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Printf("s4 pointer:%+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
