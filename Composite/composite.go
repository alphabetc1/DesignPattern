/*
Composite 组合模式
定义：将对象组合成树形结构以表示“部分-整体”的层次结构，组合模式使得用户对单个对象和组合对象的使用具有一致性。
作者：   alphabetc1
日期：   20210219
*/

package main

import (
	"fmt"
	"strings"
)

type Company interface {
	Add(Company)
	Remove(Company)
	Display(int)
	LineOfDuty()
}

type ConcreteCompany struct {
	name string
	list []Company
}

func NewConcreateCompany(name string) *ConcreteCompany {
	return &ConcreteCompany{name, []Company{}}
}

func (c *ConcreteCompany) Add(company Company) {
	if c == nil {
		return
	}

	c.list = append(c.list, company)
}

func (c *ConcreteCompany) Remove(company Company) {
	if c == nil {
		return
	}

	for i, val := range c.list {
		if val == company {
			c.list = append(c.list[:i], c.list[i+1:]...)
			return
		}
	}
}

func (c *ConcreteCompany) Display(depth int) {
	if c == nil {
		return
	}

	fmt.Println(strings.Repeat("-", depth), c.name)
	for _, val := range c.list {
		val.Display(depth + 2)
	}
}

func (c *ConcreteCompany) LineOfDuty() {
	if c == nil {
		return
	}

	for _, val := range c.list {
		val.LineOfDuty()
	}
}

type HRDepartment struct {
	name string
}

func NewHRDepartment(name string) *HRDepartment {
	return &HRDepartment{name}
}

func (h *HRDepartment) Add(c Company) {
	fmt.Println("HRDepartment Cannot Add")
}

func (h *HRDepartment) Remove(c Company) {
	fmt.Println("HRDepartment Cannot Remove")
}

func (h *HRDepartment) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth), h.name)
}

func (h *HRDepartment) LineOfDuty() {
	fmt.Println("HR", h.name, "正在监督员工!")
}

type FinanceDepartment struct {
	name string
}

func NewFinanceDepartment(name string) *FinanceDepartment {
	return &FinanceDepartment{name}
}

func (f *FinanceDepartment) Add(c Company) {
	fmt.Println("HRDepartment Cannot Add")
}

func (f *FinanceDepartment) Remove(c Company) {
	fmt.Println("HRDepartment Cannot Remove")
}

func (f *FinanceDepartment) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth), f.name)
}

func (f *FinanceDepartment) LineOfDuty() {
	fmt.Println("财务", f.name, "正在管理财务!")
}
