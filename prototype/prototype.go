/*
prototype原型模式
定义：用原型实例制定创建对象的种类，并通过拷贝这些原型创建新的对象
作者：   alphabet
日期：   20201217
*/
package main

import "fmt"

type IClonable interface {
	Clone()
}

type WorkExperience struct {
	workDate string
	company  string
}

func (w *WorkExperience) Clone() *WorkExperience {
	if w == nil {
		return nil
	}
	newWorkExperience := (*w)
	return &newWorkExperience
}

func (w *WorkExperience) CloneWithNewExperience(_workDate, _company string) *WorkExperience {
	if w == nil {
		return nil
	}

	return &WorkExperience{
		workDate: _workDate,
		company:  _company,
	}
}

type Resume struct {
	name string
	sex  string
	age  string
	work WorkExperience
}

func (r *Resume) SetPersonalInfo(name, sex, age string) {
	if r == nil {
		return
	}
	r.name = name
	r.sex = sex
	r.age = age
}

func (r *Resume) SetWorkExperience(workDate, company string) {
	if r == nil {
		return
	}
	r.work.workDate = workDate
	r.work.company = company
}

func (r *Resume) Display() {
	if r == nil {
		return
	}
	fmt.Println("个人信息：", r.name, r.sex, r.age)
	fmt.Println("工作经历：", r.work.workDate, r.work.company)
}

func (r *Resume) Clone() *Resume {
	if r == nil {
		return nil
	}
	return &Resume{
		name: r.name,
		age:  r.age,
		sex:  r.sex,
		work: *(r.work.Clone()),
	}
}

func (r *Resume) CloneWithNewExperience(_workDate, _company string) *Resume {
	if r == nil {
		return nil
	}
	return &Resume{
		name: r.name,
		age:  r.age,
		sex:  r.sex,
		work: *(r.work.CloneWithNewExperience(_workDate, _company)),
	}
}

func NewResume() *Resume {
	return &Resume{}
}
