package base

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// 问题一
// 字符串转成数组会发生拷贝吗？
func TestCase1(t *testing.T) {
	/*
			在go里，字符串是不可变的，一旦创建就不能修改
		而[]byte是可以修改，那把一个不可变的类型转为一个可修改的类型肯定是要发生拷贝的
	*/

	/**
	字符串到字节数组的转换涉及到内存分配和数据复制，因此在性能敏感的场景中需要考虑其影响。

	Unicode 处理：对于包含 Unicode 字符的字符串，转换为字节数组时需要注意编码问题，确保不会丢失或错误处理字符。
	*/
}

// 有办法不拷贝吗？
func TestName(t *testing.T) {
	a := "aaa"
	// StringHeader 是string的底层结构
	/*
		type StringHeader struct {
			Data uintptr
			Len  int
		}
	*/
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&a))

	var byteslice []byte
	// SliceHeader 是slice的底层结构
	/*
		type SliceHeader struct {
			Data uintptr
			Len  int
			Cap  int
		}
	*/
	bytesliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&byteslice))

	// 转换
	bytesliceHeader.Data = strHeader.Data
	bytesliceHeader.Len = strHeader.Len
	bytesliceHeader.Cap = strHeader.Len

	fmt.Printf("%v", byteslice)
}
