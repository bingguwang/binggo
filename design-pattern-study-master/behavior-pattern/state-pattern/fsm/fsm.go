package fsm

import (
	"design-pattern-study-master/behavior-pattern/state-pattern/event"
	"design-pattern-study-master/behavior-pattern/state-pattern/status"
	"fmt"
	"log"
	"sync"
)

// FSMHandler 事件处理器, 触发事件时要做的操作
type FSMHandler func(status FSMStatus, event FSMEvent) FSMStatus // 传入状态和事件返回一个新事件

// 状态机具体实现类
type fsm struct {
	mu       sync.Mutex                 // 排它锁
	State    FSMStatus                  // call handler之后的状态，也就是计算后输出的状态
	Handlers map[int]map[int]FSMHandler // 状态-事件-处理器map
}

// AddHandler 组装状态机，把handler根据 状态status和触发事件event 装配到状态机的handlers表里
func (f *fsm) AddHandler(status FSMStatus, event FSMEvent, handler FSMHandler) FSM {
	if _, ok := f.Handlers[status.FSMStatusID()]; !ok {
		f.Handlers[status.FSMStatusID()] = make(map[int]FSMHandler)
	}
	if _, ok := f.Handlers[status.FSMStatusID()][event.FSMEventID()]; !ok {
		f.Handlers[status.FSMStatusID()][event.FSMEventID()] = handler
	}
	return f
}

// Call 执行状态机 ，根据传入的当前状态和事件去handlers表找到handler，执行handler
func (f *fsm) Call(status FSMStatus, event FSMEvent) FSMStatus {
	if _, ok := f.Handlers[status.FSMStatusID()]; !ok {
		log.Fatal("没有状态的映射关系")
	}
	if _, ok := f.Handlers[status.FSMStatusID()][event.FSMEventID()]; !ok {
		log.Fatal("没有事件的映射关系")
	}
	handler := f.Handlers[status.FSMStatusID()][event.FSMEventID()]
	f.State = handler(status, event)
	return f.State
}

var (
	once        sync.Once
	fsmInstance FSM
)

// NewFSM 定制fsm,自定义handler，并组装入fsm
func NewFSM() FSM {
	once.Do(func() {
		fsmInstance = &fsm{
			State:    nil,
			Handlers: make(map[int]map[int]FSMHandler),
		}
		fsmInstance.AddHandler(
			status.ModelStatusToPay, // 未支付状态
			event.PayOrder,          // 支付事件
			func(s FSMStatus, e FSMEvent) FSMStatus {
				// 可以在这里新增相关业务处理逻辑,比如发送mq等工作，有限状态机本质上还是个条件判断，只是通过枚举的方式把它实现了而已
				fmt.Println("dome something ...")
				return status.ModelStatusToDeliver
			},
		).AddHandler(
			status.ModelStatusToDeliver, // 待发货状态
			event.DeliverOrder,          // 发货事件
			func(s FSMStatus, e FSMEvent) FSMStatus {
				// 可以在这里新增相关业务处理逻辑
				fmt.Println("dome something ...")
				return status.ModelStatusDelivered
			},
		).AddHandler(
			status.ModelStatusDelivered, // 已发货状态
			event.ReceiveOrder,          // 签收事件
			func(s FSMStatus, e FSMEvent) FSMStatus {
				// 可以在这里新增相关业务处理逻辑
				fmt.Println("dome something ...")
				return status.ModelStatusReceivered
			},
		)
	})
	return fsmInstance
}
