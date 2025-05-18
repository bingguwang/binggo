package fsm

// FSMStatus 状态机内的状态
type FSMStatus interface {
    FSMStatusID() int
    FSMStatusDesc() string
}

// FSMEvent 状态机内的触发事件
type FSMEvent interface {
    FSMEventID() int
    FSMEventDesc() string
}

// FSM 状态机接口
type FSM interface {
    AddHandler(status FSMStatus, event FSMEvent, handler FSMHandler) FSM
    Call(status FSMStatus, event FSMEvent) FSMStatus
}

/**
  当前状态 + 触发事件 ==》状态机==》get handler(要如何处理) ==》 call handler 实现状态的转变
*/
