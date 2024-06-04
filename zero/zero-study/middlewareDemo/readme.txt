在 go-zero 中使用 HTTP Middleware 共有两种方式。

全局 Middleware 配置
直接在main里加上	server.Use(middlewareDemoFunc) // 在这加的是全局中间件


路由组 Middleware 配置
这类可以先在api文件里写名，生成代码后去实现相应的中间件
@server(
    middleware: GreetMiddleware1, GreetMiddleware2
)

测试下可以看到，最后middle执行顺序是，全局中间件的前置逻辑--GreetMiddleware1的前置逻辑--GreetMiddleware2的前置逻辑--处理--
    --GreetMiddleware2的后置逻辑--GreetMiddleware1的后置逻辑--全局中间件的后置逻辑
