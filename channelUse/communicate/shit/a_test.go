package shit

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type gatway struct {
	oschan   chan os.Signal
	exitchan chan int
}

func newgatway(oschan chan os.Signal, exitchan chan int) *gatway {
	return &gatway{
		oschan:   oschan,
		exitchan: exitchan,
	}
}
func (g *gatway) run() {
	defer close(g.exitchan)
	taskcount := 0
	for {
		select {
		case <-g.oschan:
			fmt.Println("中断退出，网关被动退出")
			return
		case <-time.After(3 * time.Second):
			taskcount++
			fmt.Println("网关运行中")
			if taskcount == 3 {
				return
			}
		}
	}
}
func TestName(t *testing.T) {
	oschan := make(chan os.Signal, 1)
	exitchan := make(chan int, 1)
	g := newgatway(oschan, exitchan)

	g.run()

	for {
		select {
		case <-exitchan:
			fmt.Println("网关退出")
			return
		case <-oschan:
			fmt.Println("中断信号退出")
			return
		}
	}

}
