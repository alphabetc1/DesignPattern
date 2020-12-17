/*
第二种抽象类的实现方式
作者：   alphabet
日期：   20201217
*/
package main

import "fmt"

type Speak interface {
	FirstSentence()
	SecondSentence()
	ThirdSentence()
}

type Template struct{}

func (t *Template) Talk(s Speak) {
	s.FirstSentence()
	s.SecondSentence()
	s.ThirdSentence()
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
