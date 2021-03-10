package main

import "fmt"

type Request struct {
	requestType    string
	requestContent string
	number         int
}

type Manager interface {
	SetSuperior(m Manager)
	RequestApplications(request *Request)
}

type CommonManager struct {
	name     string
	superior Manager
}

func (c *CommonManager) SetSuperior(superior Manager) {
	if c == nil {
		return
	}
	c.superior = superior
}

func (c *CommonManager) RequestApplications(request *Request) {
	if request.requestType == "请假" && request.number <= 2 {
		fmt.Println(c.name, ": ", request.requestContent, "数量", request.number, "被批准")
	} else {
		if c.superior != nil {
			c.superior.RequestApplications(request)
		}
	}
}

type Majordomo struct {
	name     string
	superior Manager
}

func (m *Majordomo) SetSuperior(superior Manager) {
	if m == nil {
		return
	}
	m.superior = superior
}

func (m *Majordomo) RequestApplications(request *Request) {
	if request.requestType == "请假" && request.number <= 10 {
		fmt.Println(m.name, ": ", request.requestContent, "数量", request.number, "被批准")
	} else {
		if m.superior != nil {
			m.superior.RequestApplications(request)
		}
	}
}

type GeneralManager struct {
	name     string
	superior Manager
}

func (g *GeneralManager) SetSuperior(superior Manager) {
	if g == nil {
		return
	}
	g.superior = nil
}

func (g *GeneralManager) RequestApplications(request *Request) {
	if request.requestType == "请假" {
		fmt.Println(g.name, ": ", request.requestContent, "数量", request.number, "被批准")
	} else if request.requestType == "加薪" && request.number <= 500 {
		fmt.Println(g.name, ": ", request.requestContent, "数量", request.number, "被批准")
	} else {
		fmt.Println(g.name, ": ", request.requestContent, "数量", request.number, "再说吧")
	}
}
