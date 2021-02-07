/*
Builder 建造者模式
定义：将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示
作者：   alphabet
日期：   20210207
*/

package main

import "fmt"

type PersonBuilder interface {
	SpeakName()
	BuildHead()
	BuildBody()
	BuildArm()
	BuildLeg()
}

type PersonThin struct {
	name string
}

func (p *PersonThin) SpeakName() {
	fmt.Println("俺叫", p.name)
}

func (p *PersonThin) BuildHead() {
	fmt.Println("俺的脑袋瘦")
}

func (p *PersonThin) BuildBody() {
	fmt.Println("俺的身体瘦")
}

func (p *PersonThin) BuildArm() {
	fmt.Println("俺的胳膊瘦")
}

func (p *PersonThin) BuildLeg() {
	fmt.Println("俺的大腿瘦")
}

type PersonFat struct {
	name string
}

func (p *PersonFat) SpeakName() {
	fmt.Println("俺叫", p.name)
}

func (p *PersonFat) BuildHead() {
	fmt.Println("俺的脑袋胖")
}

func (p *PersonFat) BuildBody() {
	fmt.Println("俺的身体胖")
}

func (p *PersonFat) BuildArm() {
	fmt.Println("俺的胳膊胖")
}

func (p *PersonFat) BuildLeg() {
	fmt.Println("俺的大腿胖")
}

type PersonDirector struct {
	person PersonBuilder
}

func (p *PersonDirector) CreatePerson() {
	if p == nil {
		return
	}
	p.person.SpeakName()
	p.person.BuildHead()
	p.person.BuildBody()
	p.person.BuildArm()
	p.person.BuildLeg()
}
