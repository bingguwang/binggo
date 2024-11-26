package main

import (
	"sync"
	"testing"
)

/*
*
嵌套锁定：如果你的代码中有其他函数也锁定了同一个互斥锁，并且这些函数相互调用，可能会导致死锁。
例如，如果某个其他函数在持有锁的情况下调用 UpdateDuration，那么会发生死锁。

重复锁定：如果在同一个代码路径中，多次尝试锁定相同的互斥锁，并且这些锁定操作没有解锁（即：没有调用Unlock），则会导致死锁。
例如，在一些错误处理代码中，可能会错误地重复调用 Lock。
*/
func TestCase1(t *testing.T) {
	s := &MyS{}
	s.other()
}

type MyS struct {
	lock sync.Mutex
}

func (s *MyS) UpdateDuration() {
	s.lock.Lock()
	defer s.lock.Unlock()
}

func (s *MyS) other() {
	s.lock.Lock()
	defer s.lock.Unlock()

	//多次尝试锁定相同的互斥锁，调用UpdateDuration内会锁住，但是此时锁并未先解锁
	s.UpdateDuration() // fatal error: all goroutines are asleep - deadlock!
}

/**
确保锁只被一个 Goroutine 持有。不要在持有锁的情况下调用可能会锁定同一个互斥锁的函数。
尽量减少锁定的范围。只在需要保护共享资源的代码段中锁定。
避免嵌套锁定。如果非要嵌套锁定不同的互斥锁，确保遵循一致的锁定顺序。
*/
