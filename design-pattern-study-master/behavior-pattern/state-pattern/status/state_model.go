package status

// Status 状态结构
type Status struct {
    StatusID   int
    StatusDesc string
}

// ModelStatus 状态机内的状态结构
// 有三类状态，订单状态,配送状态,送货状态, 把他们进一步封装到一个状态里，用于作为状态机的输入
// 因为要输入到状态机里，所以需要符合状态机内状态标准，所以需要实现 FSMStatus 接口
type ModelStatus struct {
    TotalStatus   Status
    OrderStatus   Status
    DeliverStatus Status
}

// 状态枚举
var (
    // 未支付
    ModelStatusToPay = &ModelStatus{
        TotalStatus:   Status{StatusID: TOTAL_STATUS_TO_PAY_FLAG, StatusDesc: TOTAL_STATUS_MAP[TOTAL_STATUS_TO_PAY_FLAG]},
        OrderStatus:   Status{StatusID: ORDER_STATUS_TO_PAY_FLAG, StatusDesc: ORDER_STATUS_MAP[ORDER_STATUS_TO_PAY_FLAG]},
        DeliverStatus: Status{StatusID: DELIVER_STATUS_DONT_DELIVER_FLAG, StatusDesc: DELIVER_STATUS_DELIVER_MAP[DELIVER_STATUS_DONT_DELIVER_FLAG]},
    }

    // 待发货
    ModelStatusToDeliver = &ModelStatus{
        TotalStatus:   Status{StatusID: TOTAL_STATUS_TO_DELIVER_FLAG, StatusDesc: TOTAL_STATUS_MAP[TOTAL_STATUS_TO_DELIVER_FLAG]},
        OrderStatus:   Status{StatusID: ORDER_STATUS_TO_DELIVER_FLAG, StatusDesc: ORDER_STATUS_MAP[ORDER_STATUS_TO_DELIVER_FLAG]},
        DeliverStatus: Status{StatusID: DELIVER_STATUS_TO_DELIVER_FLAG, StatusDesc: DELIVER_STATUS_DELIVER_MAP[DELIVER_STATUS_TO_DELIVER_FLAG]},
    }

    // 已发货
    ModelStatusDelivered = &ModelStatus{
        TotalStatus:   Status{StatusID: TOTAL_STATUS_TO_RECEIVE_FLAG, StatusDesc: TOTAL_STATUS_MAP[TOTAL_STATUS_TO_RECEIVE_FLAG]},
        OrderStatus:   Status{StatusID: ORDER_STATUS_TO_RECEIVE_FLAG, StatusDesc: ORDER_STATUS_MAP[ORDER_STATUS_TO_RECEIVE_FLAG]},
        DeliverStatus: Status{StatusID: DELIVER_STATUS_DELIVERED_FLAG, StatusDesc: DELIVER_STATUS_DELIVER_MAP[DELIVER_STATUS_DELIVERED_FLAG]},
    }

    // 已签收
    ModelStatusReceivered = &ModelStatus{
        TotalStatus:   Status{StatusID: TOTAL_STATUS_TO_COMMENT_FLAG, StatusDesc: TOTAL_STATUS_MAP[TOTAL_STATUS_TO_COMMENT_FLAG]},
        OrderStatus:   Status{StatusID: ORDER_STATUS_TO_COMMENT_FLAG, StatusDesc: ORDER_STATUS_MAP[ORDER_STATUS_TO_COMMENT_FLAG]},
        DeliverStatus: Status{StatusID: DELIVER_STATUS_RECEIVED_FLAG, StatusDesc: DELIVER_STATUS_DELIVER_MAP[DELIVER_STATUS_RECEIVED_FLAG]},
    }
)

// 因为会用到FSM接口的方法，所以要实现FSMStatus接口才行
func (e *ModelStatus) FSMStatusID() int {
    return e.TotalStatus.StatusID
}

func (e *ModelStatus) FSMStatusDesc() string {
    return e.TotalStatus.StatusDesc
}
