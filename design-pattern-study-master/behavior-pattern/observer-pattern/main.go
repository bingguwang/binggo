package main

import "design-pattern-study-master/behavior-pattern/observer-pattern/model"

/*
*

	观察者模式
	    也叫发布订阅模式

	     你需要为发布者类添加订阅机制， 让每个对象都能订阅或取消订阅发布者事件流
	      你需要:
	          1.存储订阅者的一个数据结构
	          2.增删订阅者的公有方法
	          3.发布者仅通过接口和订阅者交流
*/
func main() {
	publisher := model.NewBookPublisher()
	tom := &model.BookObserver{Name: "tom"}
	mick := &model.BookObserver{Name: "mick"}
	// 订阅
	publisher.AddObserver(tom, mick)

	bookPublisher := publisher.(*model.BookPublisher)
	bookPublisher.UpdateMsg("fffffffffffffffffffk")

	publisher.NotifyAll()
}
