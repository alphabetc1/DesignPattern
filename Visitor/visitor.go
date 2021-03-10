package main

import (
	"fmt"
)

type Person interface {
	Accept(Action)
}

type Man struct {
	name string
}

func (m *Man) Accept(visitor Action) {
	visitor.GetManConclusion(m)
}

type Woman struct {
	name string
}

func (w *Woman) Accept(visitor Action) {
	visitor.GetWomanConclusion(w)
}

type Action interface {
	GetManConclusion(*Man)
	GetWomanConclusion(*Woman)
}

type Success struct {
	name string
}

func (s *Success) GetManConclusion(man *Man) {
	fmt.Println(man.name, s.name, "时，背后多半有一个伟大的女人。")
}

func (s *Success) GetWomanConclusion(women *Woman) {
	fmt.Println(women.name, s.name, "时，背后多半有一个不成功的男人。")
}

type Fail struct {
	name string
}

func (f *Fail) GetManConclusion(man *Man) {
	fmt.Println(man.name, f.name, "时，闷头喝酒，谁也不用劝")
}

func (f *Fail) GetWomanConclusion(women *Woman) {
	fmt.Println(women.name, f.name, "时，眼泪汪汪，谁也劝不了。")
}

type ObjectStructure struct {
	elements []Person
}

func (o *ObjectStructure) Attach(element Person) {
	if o == nil {
		return
	}
	o.elements = append(o.elements, element)
}

func (o *ObjectStructure) Detach(element Person) {
	if o == nil {
		return
	}
	for i, val := range o.elements {
		if val == element {
			o.elements = append(o.elements[:i], o.elements[i+1:]...)
			return
		}
	}
}

func (o *ObjectStructure) Display(visitor Action) {
	if o == nil {
		return
	}
	for _, val := range o.elements {
		val.Accept(visitor)
	}
}
