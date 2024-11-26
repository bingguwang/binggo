package main

import "testing"

func TestWorker(t *testing.T) {
	w := &workerStack{make([]worker, 10)}
	w.detach() // 获取一个worker
}

type workerStack struct {
	workerQueue []worker
}

type worker interface {
}

func (wq *workerStack) len() int {
	return len(wq.workerQueue)
}

/*
*
模拟一个pop的过程, 返回栈顶的worker
*/
func (wq *workerStack) detach() worker {
	l := wq.len()
	if l == 0 {
		return nil
	}

	w := wq.workerQueue[l-1]
	wq.workerQueue[l-1] = nil // 这样可以避免协程泄漏
	wq.workerQueue = wq.workerQueue[:l-1]

	return w
}
