
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




