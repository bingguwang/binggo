package event

// 事件枚举
const (
    EVENT_PAY_ORDER       = iota // 支付事件
    EVENT_DELIVER_ORDER          // 发货事件
    EVENT_RECEIVEED_ORDER        // 签收事件
)

// 事件描述
const (
    EVENT_PAY_ORDER_DESC       = "事件：支付订单"
    EVENT_DELIVER_ORDER_DESC   = "事件：订单发货"
    EVENT_RECEIVEED_ORDER_DESC = "事件：订单收货"
)

var ModelEventMap = map[int]string{
    EVENT_PAY_ORDER:       EVENT_PAY_ORDER_DESC,
    EVENT_DELIVER_ORDER:   EVENT_DELIVER_ORDER_DESC,
    EVENT_RECEIVEED_ORDER: EVENT_RECEIVEED_ORDER_DESC,
}
