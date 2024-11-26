package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
切片的数据结构其实是：
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
	所有传参都是值拷贝，将切片作为函数参数传递时，实质上复制的是切片的 SliceHeader，对应的底层数组是保持不变的(Data指针的值不变)

*/

func TestName1(t *testing.T) {
	s := make([]int, 0, 2)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s)))
	doSomething(s) // 拷贝的是结构SliceHeader
	fmt.Println(s)
	fmt.Println(len(s)) // 0 ，s的len和cap没变，&{824634368688 0 2}
	fmt.Println(s[:1])  // 为什么可以输出追加的元素，因为s[:1]会创建一个新的切片，也即新的SliceHeader&{824634368688 1 2}
}

func doSomething(a []int) {
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&a)))
	a = append(a, 11) // 底层数组足够，底层数组还是原来的,和s用是同一个底层数组
	//a = append(a, []int{1, 2, 3}...) // 底层数组不够了，重新分配新的底层数组，底层数组不是原来的了，和s用的不是同一个底层数组！
	fmt.Println(a[:2]) // [11,0]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&a)))
}
