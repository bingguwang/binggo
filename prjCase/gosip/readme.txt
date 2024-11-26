首先，公司里国标的这一套，底层使用的其实是下面的这个项目
https://github.com/ghettovoice/gosip
由于开源协议的问题，就把这个项目的部分内容提取出来用到项目里了而已

在这里学习一下这个gosip，以及如何借鉴此项目的

gosip的server是如下的结构:
type server struct {
	running         abool.AtomicBool
	tp              transport.Layer
	tx              transaction.Layer
	host            string
	ip              net.IP
	hwg             *sync.WaitGroup
	hmu             *sync.RWMutex
	requestHandlers map[sip.RequestMethod]RequestHandler
	extensions      []string
	userAgent       string

	log log.Logger
}

这个server结构对应了videogateway里的SipStack

而main里的 server的启动的方式:	srv.Listen("tcp", "0.0.0.0:5081")
其实对应了的videogateway里的func (gbGateway *GbGatewayManager) initGatewayStack() 方法里的 gbGateway.stack.Listen("tcp", listen)


请求如何接收的？
首先请求是在监听的部分接收的，也就是在服务启动的地方，也就是
go srv.serve()，在这个serve()里
在这里可以看到request，response，ack，err等都是在这里收到的，收到之后就开启一个协程handleRequest去处理
可以看到request接收的地方: case tx, ok := <-srv.tx.Requests():
接收request是由server的事务层接收的













































































