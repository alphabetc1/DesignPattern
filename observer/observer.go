/*
Observer 建造者模式
定义：定义一种一对多的依赖关系，让多个观察者对象同时监听某一抽象对象。
		 这个抽象对象在状态发生变化时会通知所有观察者，又叫“发布-订阅模式”
		 适用于一个对象的改变需要同时同时改变其他对象的时候
作者：   alphabet
日期：   20210208
*/
package main

import "fmt"

type Observer interface {
	Update()
}

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}

type ConerteSubject struct {
	subjectState string
	observers    []Observer
}

func (c *ConerteSubject) Attach(observer Observer) {
	if c == nil {
		return
	}
	for _, o := range c.observers {
		if o == observer {
			return
		}
	}
	c.observers = append(c.observers, observer)
}

func (c *ConerteSubject) Detach(observer Observer) {
	if c == nil {
		return
	}
	for i, o := range c.observers {
		if o == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
		}
	}
}

func (c *ConerteSubject) Notify() {
	for _, o := range c.observers {
		o.Update()
	}
}

type ConcreteObserver struct {
	name          string
	observerState string
	subject       *ConerteSubject
}

func (c *ConcreteObserver) Update() {
	fmt.Print("观察者 ", c.name, " 的旧状态是 ", c.observerState)
	c.observerState = c.subject.subjectState
	fmt.Println(" 新状态是", c.observerState)
}
