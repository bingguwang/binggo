package main

import model2 "design-pattern-study-master/behavior-pattern/chain-of-responsibility-pattern/model"

/*
*

	责任链
	      允许你将请求沿着处理者链进行发送。
	      收到请求后， 每个处理者均可对请求进行处理， 或将其传递给链上的下个处理者

	  gin的中间件就是这样的
*/
func main() {
	client := &model2.Client{}

	// 创建doctor中间件
	doctor := &model2.Doctor{}
	drugstore := &model2.Drugstore{}
	reception := &model2.Reception{}
	cashier := &model2.Cashier{}
	finish := &model2.Finish{}

	// 注册中间件
	h := UseMiddleware(doctor, reception, drugstore, cashier, finish)
	Exec(h, client)
	// 执行结果可以看到，保证中间件都执行到了， 而且执行顺序按照传入的顺序
}

func UseMiddleware(handlers ...model2.Handler) model2.Handler {
	for i := 0; i < len(handlers)-1; i++ {
		handlers[i].SetNext(handlers[i+1])
	}

	// gin里把这个handlers数组作为context的成员 handlers

	return handlers[0]
}
func Exec(handler model2.Handler, client *model2.Client) {
	handler.Handle(client)
}
