package pool

import (
	"fmt"
	"sync"
	"testing"
)

/*
如果说对象池里的对象一定不会GC，是不对的
并不保证对象会一直保留在池中，当内存不足或者触发垃圾回收时，池中的对象可能会被清理。

注意的是：池子里有多少个对象是未知的，对象被释放的时机也是未知的
由于用完的对象放回池里，所以不适合有状态的对象

网络包收取发送的时候，用 sync.Pool 会有奇效
*/
func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} { // 当池里没有可用对象时，会使用这里定义的方法新建对象
			// Sync.Pool本身是并发安全的，但是如果自定义的New函数里有并发安全问题，那就会有并发问题
			// 所以在自定义的new函数里，应该注意使用互斥锁来保证逻辑是并发安全的
			return new(int)
		},
	}

	// 从池中获取一个对象
	v := pool.Get().(*int)
	fmt.Println(*v) // 打印 0，默认值

	// 修改对象的值
	*v = 42

	// 将对象放回池中
	pool.Put(v)
	// 如果只调用 Get 不调用 Put 会怎么样？
	// 只进行 Get 操作的话，就相当于一直在生成新的对象，Pool 池也失去了它最本质的功能。

	// 再次从池中获取对象
	v2 := pool.Get().(*int)
	fmt.Println(*v2) // 打印 42，因为对象是从池中获取的
}
