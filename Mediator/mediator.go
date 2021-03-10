package main

import "fmt"

type UniteNations interface {
	Declare(string, Country)
}

type Country interface {
	Name() string
	Declare(string)
	GetMessage(string, Country)
}

type USA struct {
	name     string
	mediator UniteNations
}

func (u *USA) Name() string {
	return u.name
}

func (u *USA) Declare(message string) {
	if u == nil {
		return
	}
	u.mediator.Declare(message, u)
}

func (u *USA) GetMessage(message string, country Country) {
	if u == nil {
		return
	}
	fmt.Println(u.name, "获得联合国转发", country.Name(), "的信息信息：", message)
}

type China struct {
	name     string
	mediator UniteNations
}

func (c *China) Name() string {
	return c.name
}

func (c *China) Declare(message string) {
	if c == nil {
		return
	}
	c.mediator.Declare(message, c)
}

func (c *China) GetMessage(message string, country Country) {
	if c == nil {
		return
	}
	fmt.Println(c.name, "获得联合国转发", country.Name(), "的信息信息：", message)
}

type Europe struct {
	name     string
	mediator UniteNations
}

func (e *Europe) Name() string {
	return e.name
}

func (e *Europe) Declare(message string) {
	if e == nil {
		return
	}
	e.mediator.Declare(message, e)
}

func (e *Europe) GetMessage(message string, country Country) {
	if e == nil {
		return
	}
	fmt.Println(e.name, "获得联合国转发", country.Name(), "的信息信息：", message)
}

type UnitedNationsCouncil struct {
	name     string
	countrys []Country
}

func NewUnitedNationsCouncil(name string) *UnitedNationsCouncil {
	return &UnitedNationsCouncil{name, nil}
}

func (u *UnitedNationsCouncil) AddCountry(country Country) {
	if u == nil {
		return
	}
	u.countrys = append(u.countrys, country)
}

func (u *UnitedNationsCouncil) Declare(message string, country Country) {
	if u == nil {
		return
	}
	for _, val := range u.countrys {
		if val == country {
			continue
		}
		val.GetMessage(message, country)
	}
}
