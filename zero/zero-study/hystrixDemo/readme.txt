熔断和降级

熔断是调用方的一种措施
而降级是被调用端，服务端的一种自我保护机制

熔断器一般有三种状态
1.关闭：默认状态，请求能被到达目标服务，同时统计在窗口时间成功和失败次数，如果达到错误率阈值将会进入断开状态。
2.断开： 此状态下将会直接返回错误，如果有 fallback 配置则直接调用 fallback 方法。
3.半断开：进行断开状态会维护一个超时时间（此时间内由于是断开状态可以有效缓解服务端压力），到达超时时间开始进入 半断开 状态，
尝试允许一部门请求正常通过并统计成功数量，
判断：成功次数足够 -> 转为关闭状态
        失败次数足够 -> 转为断开状态
如果这些试探性请求大多数成功，熔断器会恢复到关闭状态，否则进入 断开 状态，继续拒绝所有请求。

半断开 状态存在的目的在于实现了自我修复，同时防止正在恢复的服务再次被大量打垮。

简单总结就是：
。当请求失败比率达到一定阈值之后，熔断器开启，并休眠一段时间（由配置决定），这段休眠期过后，熔断器将处于半开状态，
在此状态下将试探性的放过一部分流量，如果这部分流量调用成功后，再次将熔断器关闭，否则熔断器继续保持开启并进入下一轮休眠周期。


影响熔断器判断的参数有很多
错误比例阈值：达到该阈值进入 断开 状态。
断开状态超时时间:超时后进入 半断开 状态。
半断开状态允许请求数量。
窗口时间大小。

而一般我们不太了解这些怎么设置比较合理
然后就有了一系列的自适应熔断器
常见的有hystrix-go、zero有breaker

这里我们看下breaker
https://github.com/zeromicro/zero-doc/blob/main/docs/zero/%E8%87%AA%E9%80%82%E5%BA%94%E8%9E%8D%E6%96%AD%E4%B8%AD%E9%97%B4%E4%BB%B6.md

注意！！！！！！！！！！！

breaker以被内置到zero框架内了，不需要显式去对请求进行熔断处理，已经内置了
HTTP 以请求方法+路由作为统计维度
用 HTTP 状态码 500 作为错误采集指标进行统计，
详情可参考 breakerhandler.go

gRPC 客户端以 rpc 方法名作为统计维度，
用 grpc 的错误码为 codes.DeadlineExceeded, codes.Internal, codes.Unavailable, codes.DataLoss, codes.Unimplemented 作为错误采集指标进行统计
详情可参考 breakerinterceptor.go
写在一个拦截器里：
func BreakerInterceptor(ctx context.Context, method string, req, reply any,
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	breakerName := path.Join(cc.Target(), method)
	return breaker.DoWithAcceptableCtx(ctx, breakerName, func() error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}, codes.Acceptable)
}
看下Acceptable可以看到，错误采集指标
func Acceptable(err error) bool {
	switch status.Code(err) {
	case codes.DeadlineExceeded, codes.Internal, codes.Unavailable, codes.DataLoss,
		codes.Unimplemented, codes.ResourceExhausted:
		return false
	default:
		return true
	}
}


// todo 说说降级
