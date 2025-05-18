package model

// 让标准插座（适配器）去适配日式插座,也就是实现日式标准接口

type Adaptor struct {
    c *ChinesePort // 封装了服务对象
}

// CommonPlug 适配器接口，提供给客户调用
func (a *Adaptor) CommonPlug() {
    // 适配
    a.c.WirelessPlug()
}
