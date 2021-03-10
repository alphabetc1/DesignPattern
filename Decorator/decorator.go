package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) Show() {
	if p == nil {
		return
	}
	fmt.Println("装扮的", p.name)
}

type AbstractPerson interface {
	Show()
}

type Decorator struct {
	component AbstractPerson
}

func (d *Decorator) SetDecorator(abstractPerson AbstractPerson) {
	if d == nil {
		return
	}
	d.component = abstractPerson
}

func (d *Decorator) Show() {
	if d == nil {
		return
	}

	if d.component != nil {
		d.component.Show()
	}
}

type Tshirt struct {
	Decorator
}

func (t *Tshirt) Show() {
	if t == nil {
		return
	}
	fmt.Print("大T恤 ")
	t.Decorator.Show()
}

type BigTrouser struct {
	Decorator
}

func (b *BigTrouser) Show() {
	if b == nil {
		return
	}
	fmt.Print("垮裤 ")
	b.Decorator.Show()
}

type Suit struct {
	Decorator
}

func (s *Suit) Show() {
	if s == nil {
		return
	}
	fmt.Print("西装 ")
	s.Decorator.Show()
}

type Miniskirt struct {
	Decorator
}

func (m *Miniskirt) Show() {
	if m == nil {
		return
	}
	fmt.Print("超短裙 ")
	m.Decorator.Show()
}
