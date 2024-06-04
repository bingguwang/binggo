这个是通过一个短连接的项目来帮助学习理解zero 的rpc的使用
短连接的项目指的是一个调用链路比较简单的小demo，只有rest调rpc，mysql以及缓存的使用


首先要安装goctl, protoc , protoc-gen-go

安装goctl
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest

安装protoc , protoc-gen-go
参考这篇如何下载：https://go-zero.dev/docs/tasks/installation/protoc

（当然你可以直接使用goctl来安装protoc相关的组件，直接执行goctl env install --verbose --force）


检查安装是否成功
goctl env check --verbose


mkdir -p myprj/api

cd myprj 然后go mod init


然后进入api目录，编写api文件
编写完api文件后，生成项目代码:
goctl api go -api my.api -dir .
go mod tidy

运行sever:
go run shorturl.go -f etc/shorturl-api.yaml

到这里，已经有了rest服务接口了。
我们知道，在项目里rest接口的handler里的逻辑处理往往会用到rpc调用，这里还没有，加上rpc服务去看看

cd myprj
mkdir -p rpc/transform

进入transform目录，编写proto文件
也可以直接goctl生成:goctl rpc -o transform.proto，然后按自己需要改写


接下来就是用proto生成rpc代码了，goctl可以很方便的生成，在transform路径下：
goctl rpc protoc transform.proto --go_out=. --go-grpc_out=. --zrpc_out=.

生成的rpc目录如下:
rpc/transform
├── etc
│   └── transform.yaml              // 配置文件
├── internal
│   ├── config
│   │   └── config.go               // 配置定义
│   ├── logic
│   │   ├── expandlogic.go          // expand 业务逻辑在这里实现
│   │   └── shortenlogic.go         // shorten 业务逻辑在这里实现
│   ├── server
│   │   └── transformerserver.go    // 调用入口, 不需要修改
│   └── svc
│       └── servicecontext.go       // 定义 ServiceContext，传递依赖
├── transform
│        ├── transform.pb.go
│        └── transform_grpc.pb.go
├── transform.go                    // rpc 服务 main 函数
├── transform.proto
└── transformer
    └── transformer.go              // 提供了外部调用方法，无需修改

执行go mod tidy调整一下依赖

可以看看生成的rpc的配置文件，里面已经把etcd相关的配置都生成好了，可以根据实际情况去改写
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: transform.rpc
改为
Etcd:
  Hosts:
    - 192.168.0.58:2379
  Key: transform.rpc
  User: root
  Pass: "123456"

配置的字段名称是有规定的，可以去RpcServerConf结构里看到响应的字段

然后可以运行一下rpc服务
go run transform.go -f etc/transform.yaml
看看服务是否被注册进了etcd
可以看到服务已注册：{transform.rpc/4944266073912622447 : 192.168.2.26:8080}

然后需要去api的处理逻辑里调用rpc试试
首先api得配置一下rpc服务的配置，api的配置文件里加上
Transform:
  Etcd:
    Hosts:
      - 192.168.0.58:2379
    Key: transform.rpc

还得在Config配置定义结构里加上rpc相关的属性
rpc对应的属性类型是啥呢？
在rpc/transform的config里可以看到配置是zrpc.RpcServerConf类型，所以我们去api的config里也是添加此类型

不仅如此，ServiceContext也要去添加

至此，就可以去api的处理逻辑调用了

添加完调用后测试一下，把rest和rpc服务都开启来，调用rest接口试下:
curl -i "http://localhost:8888/shorten?url=https://www.baidu.com"

到这步可以看到调用成功了

还可以加入数据库和缓存
来试试
mkdir -p rpc/transform/model

如果想直接用zero的orm，那么对于数据相关的代码goctl也是可以生成的，生成是工具sql文件来生成
goctl model mysql ddl -c -src shorturl.sql -dir .
生成的代码如下:
rpc/transform/model
├── shorturl.sql
├── shorturlmodel.go              // 扩展代码
├── shorturlmodel_gen.go          // CRUD+cache 代码
└── vars.go                       // 定义常量和变量

连缓存的代码也生成了

然后需要去添加配置以及设置配置类
type Config struct {
	zrpc.RpcServerConf
	DataSource string          // 手动代码
	Table      string          // 手动代码
	Cache      cache.CacheConf // 手动代码
}
还需要修改servicecontext代码:
type ServiceContext struct {
	Config config.Config
	Model  model.ShorturlModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}

然后取处理逻辑里调用试试看
测试一下:
curl -i "http://localhost:8888/shorten?url=https://www.baidu.com"


注意的是：一般微服务里，每个微服务是有api也有rpc的!