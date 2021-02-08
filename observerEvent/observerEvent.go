/*
ObserverEvent 观察者模式-事件委托
定义：定义一种一对多的依赖关系，让多个观察者对象同时监听某一抽象对象。
		 这个抽象对象在状态发生变化时会通知所有观察者，又叫“发布-订阅模式”
		 适用于一个对象的改变需要同时同时改变其他对象的时候
作者：   alphabet
日期：   20210208
*/
package main

import "fmt"

type update func(string)

type Subject interface {
	Notify()
	AddEventHandler(*update)
	RemoveEventHandler(*update)
}

type Boss struct {
	action       string
	eventHandler []*update
}

func (b *Boss) Notify() {
	if b == nil {
		return
	}
	for _, e := range b.eventHandler {
		(*e)(b.action)
	}
}

func (b *Boss) AddEventHandler(event *update) {
	if b == nil {
		return
	}
	for _, e := range b.eventHandler {
		if e == event {
			return
		}
	}

	b.eventHandler = append(b.eventHandler, event)
}

func (b *Boss) RemoveEventHandler(event *update) {
	if b == nil {
		return
	}
	for i, e := range b.eventHandler {
		if e == event {
			b.eventHandler = append(b.eventHandler[:i], b.eventHandler[i+1:]...)
		}
	}
}

type StockObserver struct {
	name string
}

func (s *StockObserver) CloseStockMarket(action string) {
	fmt.Println(action, s.name, "关闭股票行情 ，继续工作！")
}

type NBAObserver struct {
	name string
}

func (n *NBAObserver) CloseNBADirectSeeding(action string) {
	fmt.Println(action, n.name, "关闭NBA直播 ，继续工作！")
}
