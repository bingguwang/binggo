
拦截器

在grpc里不陌生，有4类
客户端有2类
unary rpc对应有unary client interceptor
stream rpc对应有stream client interceptor

服务端有2类
unary rpc对应有unary server interceptor
stream rpc对应有stream server interceptor

client ---> unary rpc ---> unary client interceptor ---> unary server interceptor ---> unary handler
client ---> stream rpc ---> stream client interceptor ---> stream server interceptor ---> stream handler


1.客户端拦截器
需要实现grpc.UnaryClientInterceptor
然后再在创建客户端的时候通过zrpc.WithUnaryClientInterceptor进行注册

2.服务端拦截器
服务端拦截器需要实现grpc.UnaryServerInterceptor
通过RpcServer.AddUnaryInterceptors进行注册

这里使用客户端拦截器输出请求方法耗时，
使用服务端拦截器进行简单的限流

看下客户端拦截器怎么创建，自然是要去实例化客户端的地方创建
实例化client的地方这里是在api的NewServiceContext方法内
加上拦截器
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Bing:   bingclient.NewBing(zrpc.MustNewClient(c.Bing, zrpc.WithUnaryClientInterceptor(myTimeInterceptor))),
	}
}
myTimeInterceptor自己实现就行了

看下服务端拦截器怎么创建，同理要去实例化的服务端的地方创建
实例化server的地方在rpc的main方法部分: s := zrpc.MustNewServer
添加拦截器：s.AddUnaryInterceptors(rateLimitInterceptor)
拦截器自己实现

可以测试下限流器拦截器是否有用，去调用的地方加个循环，看看rpc server拦截器是否拦截生效


