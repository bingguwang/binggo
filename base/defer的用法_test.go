package base

import "testing"

/**
defer后面的函数只有在当前函数执行完毕后才能执行，
将延迟的语句按defer的逆序进行执行，
也就是说先被defer的语句最后被执行，最后被defer的语句，最先被执行，通常用于释放资源。

*/

/*
*
defer执行的时机：
*/
func TestCase11(t *testing.T) {

}
