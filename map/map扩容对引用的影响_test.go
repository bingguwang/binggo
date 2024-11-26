package _map

import "testing"

import "fmt"

/*
由下结果可以看到map的扩容不影响map的引用
*/
func TestName(t *testing.T) {

	myMap := make(map[string]int, 2)
	myMap["key1"] = 1
	myMap["key2"] = 2

	fmt.Printf("Before adding elements: %p\n", myMap)

	// Adding elements that may cause map to resize
	for i := 0; i < 10000; i++ {
		myMap[fmt.Sprintf("key%d", i+3)] = i + 3
	}

	fmt.Printf("After adding elements: %p\n", myMap)
}

/**
map在扩容时候，虽然可能底层的哈希表结果会发生变化，但是这不影响引用
在 Go 中，map 的底层数据结构是由运行时系统管理的，虽然它可能因为元素的插入、删除等操作而改变，
但是对于程序员来说，通过引用操作 map 时，不需要关心其底层的数据结构细节，也不需要特别考虑是否需要传递指针。
因此，函数间传递 map 时，仍然不需要使用 & 符号获取其指针。
*/
