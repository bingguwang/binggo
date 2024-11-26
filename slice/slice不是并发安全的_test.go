package main

import (
	"fmt"
	"testing"
)

/*
从下面的例子可以很直观看到，slice不是并发安全的
*/
func TestName(t *testing.T) {
	var res []int
	for i := 0; i < 50; i++ {
		go func() {
			res = append(res, i)
		}()

		// 一样的还是非并发安全的
		//go func(i int) {
		//	res = append(res, i)
		//}(i)
	}
	fmt.Println(len(res))
	// 可以看到很多的重复值，因为好几个协程读取到的下标一样也就出现了的重复值
	fmt.Println(res)
}
