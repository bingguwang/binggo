package ant

import (
	"fmt"
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	pool := NewPool(5)

	for i := 0; i < 10; i++ {
		i := i
		pool.Submit(func() {
			fmt.Println("执行任务:", i)
		})
	}
	pool.ShoutDown()
	fmt.Println("池子关闭")
}

func (p *Pool) ShoutDown() {
	p.wg.Wait()
	close(p.workers)
}

type Pool struct {
	current int
	maxsize int
	stop    chan struct{}
	workers chan task
	wg      sync.WaitGroup
	lock    sync.Mutex
}

type task func()

func NewPool(max int) *Pool {
	res := &Pool{
		current: 0,
		maxsize: max,
		workers: make(chan task, max),
	}
	return res
}

func (p *Pool) Submit(task task) {
	p.lock.Lock()
	if p.maxsize > p.current {
		p.workers <- task
		go p.run()
		p.current++
	}
	p.lock.Unlock()

}

func (p *Pool) run() {
	defer p.wg.Done()
	for task := range p.workers {
		p.wg.Add(1)
		task()
		p.lock.Lock()
		p.current--
		p.lock.Unlock()
	}
}
