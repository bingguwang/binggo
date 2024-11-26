package base

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestBig(t *testing.T) {
	fmt.Println(isLittleEndian())
}

// 判断是大端还是小端
func isLittleEndian() bool {
	var i int32 = 0x01020304
	p := *(*byte)(unsafe.Pointer(&i))
	return p == 0x04
}
