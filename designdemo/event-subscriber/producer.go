package main

type Producer interface {
	subscribe(consumer Consumer)   // 订阅
	resubscribe(consumer Consumer) // 取消订阅
	publish()                      // 发布
}
