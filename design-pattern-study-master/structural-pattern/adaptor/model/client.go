package model

type Client struct {
}

// Charge 客户端接口
func (*Client) Charge(standard Standard) {
    standard.CommonPlug() // 调用适配器接口
}
