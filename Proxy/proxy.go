package main

import "fmt"

type SchoolGirl struct {
	name string
}

func (s *SchoolGirl) Name() string {
	return s.name
}

type GiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}

type Pursuit struct {
	name string
	mm   SchoolGirl
}

func (p *Pursuit) GiveDolls() {
	fmt.Println(p.mm.Name(), p.name, "送你洋娃娃")
}

func (p *Pursuit) GiveFlowers() {
	fmt.Println(p.mm.Name(), p.name, "送你鲜花")
}

func (p *Pursuit) GiveChocolate() {
	fmt.Println(p.mm.Name(), p.name, "送你巧克力")
}

type Proxy struct {
	name string
	gg   Pursuit
}

func (p *Proxy) GiveDolls() {
	p.gg.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	p.gg.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.gg.GiveChocolate()
}
