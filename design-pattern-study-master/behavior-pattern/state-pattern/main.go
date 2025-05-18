package main

import (
	"design-pattern-study-master/behavior-pattern/state-pattern/event"
	"design-pattern-study-master/behavior-pattern/state-pattern/fsm"
	"design-pattern-study-master/behavior-pattern/state-pattern/status"
	"design-pattern-study-master/utils"
	"fmt"
)

/**
  有限状态机FSM

    以前对于状态转换，我们是直接用条件语句去实现，当状态和时间多的时候会特别难维护

    状态机相当于是一种枚举，把状态和事件都列举出来，我们只要那这当前状态和触发事件输入到状态机，状态机就会设置好下一个状态

    当前状态 + 触发事件 ==》状态机==》get handler(要如何处理) ==》 call handler 实现状态的转变

    决定了映射的其实是在handler里
*/

func main() {
	fsm := fsm.NewFSM()

	curStatus := status.ModelStatusToDeliver
	fmt.Println("当前状态: ", utils.ToJson(curStatus))

	curEvent := event.DeliverOrder
	fmt.Println("当前事件: ", utils.ToJson(curEvent))

	res := fsm.Call(curStatus, curEvent)
	fmt.Println("结果: ", utils.ToJson(res))

}
