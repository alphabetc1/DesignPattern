/*
Adapter 适配器模式
定义：将一个类的接口转换成客户希望的另外一个接口。
Adapter使得原本由于接口不兼容而不能一起工作的类可以一起工作。
适用场景：两个类功能相似但是接口不同，而且双方都不太容易修改时。
作者：   alphabetc1
日期：   20210217
*/
package main

import "fmt"

type Player interface {
	Attack()
	Defence()
}

type ForeignCenter struct {
	name string
}

func (f *ForeignCenter) Gan() {
	fmt.Println(f.name, "干就完了，奥利给！")
}

func (f *ForeignCenter) Liu() {
	fmt.Println(f.name, "溜了溜了")
}

type Translator struct {
	f *ForeignCenter
}

func (t *Translator) Attack() {
	t.f.Gan()
}

func (t *Translator) Defence() {
	t.f.Liu()
}
