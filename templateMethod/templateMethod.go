/*
template模板模式
定义：定义一个操作中的算法的骨架，而将一些具体步骤延迟到子类中。
作者：   alphabetc1
日期：   20201217
*/
package main

import "fmt"

type Speak interface {
	FirstSentence()
	SecondSentence()
	ThirdSentence()
}

type Template struct {
	s Speak //抽象基类，在子类中具体实现
}

func (t *Template) Talk() {
	if t == nil {
		return
	}
	t.s.FirstSentence()
	t.s.SecondSentence()
	t.s.ThirdSentence()
}

type Chicken struct {
	Template
}

func (c *Chicken) FirstSentence() {
	fmt.Println("吃饭")
}

func (c *Chicken) SecondSentence() {
	fmt.Println("呵呵")
}

func (c *Chicken) ThirdSentence() {
	fmt.Println("去洗澡")
}

type Dog struct {
	Template
}

func (d *Dog) FirstSentence() {
	fmt.Println("在吗")
}

func (d *Dog) SecondSentence() {
	fmt.Println("多喝热水")
}

func (d *Dog) ThirdSentence() {
	fmt.Println("早点睡")
}
