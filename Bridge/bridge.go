package main

import "fmt"

type SoftWare interface {
	Run()
}

type Game struct {
	name string
}

func (g *Game) Run() {
	fmt.Println("Software: game", g.name)
}

type Office struct {
	name string
}

func (o *Office) Run() {
	fmt.Println("Software: office", o.name)
}

type Computer struct {
	softWare SoftWare
	name     string
}

func (c *Computer) setSoftware(s SoftWare) {
	if c == nil {
		return
	}
	c.softWare = s
}

func (c *Computer) Run() {
	if c == nil {
		return
	}
	fmt.Println("Computer ", c.name, "is running")
	c.softWare.Run()
}

type ThinkPad struct {
	Computer
}

func NewThinkPad(software SoftWare, name string) *ThinkPad {
	return &ThinkPad{Computer{software, name}}
}

type Dell struct {
	Computer
}

func NewDell(software SoftWare, name string) *Dell {
	return &Dell{Computer{software, name}}
}
