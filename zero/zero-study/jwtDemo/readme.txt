
这里讲述了如何进行jwt生成token以及jwt用户身份校验

首先是定义api，定义一个login接口
然后在login的处理逻辑里，先rpc调用user服务，这个user服务也要自己实现
在user服务里取数据库里查找用户信息然后返回，然后收到rpc响应后根据用户信息生成token返回rest

先创建api目录， 在api下编写api文件并生成api代码
goctl api go -api rest.api -dir .

先创建rpc/user目录， 在user目录下编写proto文件并生成rpc代码

如何鉴权呢？
在api文件里加上如下内容作为分隔，上面的是不用鉴权的接口，下面是要鉴权的接口
@server(
	jwt: Auth
)


