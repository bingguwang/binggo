package event

// Event 事件结构
type Event struct {
    EventID   int    // 事件类别ID
    EventDesc string // 事件描述
}

// 事件实例列举
var (
    // 支付事件
    PayOrder = &Event{
        EVENT_PAY_ORDER,
        ModelEventMap[EVENT_PAY_ORDER],
    }

    // 发货事件
    DeliverOrder = &Event{
        EVENT_DELIVER_ORDER,
        ModelEventMap[EVENT_DELIVER_ORDER],
    }

    // 签收事件
    ReceiveOrder = &Event{
        EVENT_RECEIVEED_ORDER,
        ModelEventMap[EVENT_RECEIVEED_ORDER],
    }
)

// 因为要输入到状态机里，所以需要符合状态机内事件标准，所以需要实现 FSMEvent 接口

func (e *Event) FSMEventID() int {
    return e.EventID
}

func (e *Event) FSMEventDesc() string {
    return e.EventDesc
}
