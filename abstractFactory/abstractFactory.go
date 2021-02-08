/*
AbstractFactory 抽象工厂模式
定义：提供一个创建一系列相关或互相依赖对象的接口，而无需指定它们具体的类。
作者：   alphabetc1
日期：   20210208
*/
package main

import "fmt"

type AbstractFactory interface {
	CreateHuoguo() HuoguoProduct
	CreateChuanchuan() ChuanchuanProduct
}

type HuoguoProduct interface {
	HuoguoInfo()
}

type JuBaoYuan struct {
}

func (j *JuBaoYuan) HuoguoInfo() {
	fmt.Println("火锅一哥：聚宝源")
}

type YangXieZi struct {
}

func (y *YangXieZi) HuoguoInfo() {
	fmt.Println("火锅二哥：羊蝎子")
}

type ChuanchuanProduct interface {
	ChuanchuanInfo()
}

type MaLuBianBian struct {
}

func (m *MaLuBianBian) ChuanchuanInfo() {
	fmt.Println("串串一哥：马路边边")
}

type YuLin struct {
}

func (y *YuLin) ChuanchuanInfo() {
	fmt.Println("串串二哥：玉林")
}

type YigeFactory struct {
}

func (y *YigeFactory) CreateHuoguo() HuoguoProduct {
	if y == nil {
		return nil
	}
	h := &JuBaoYuan{}
	return h
}

func (y *YigeFactory) CreateChuanchuan() ChuanchuanProduct {
	if y == nil {
		return nil
	}
	c := &MaLuBianBian{}
	return c
}

type ErgeFactory struct {
}

func (e *ErgeFactory) CreateHuoguo() HuoguoProduct {
	if e == nil {
		return nil
	}
	h := &YangXieZi{}
	return h
}

func (e *ErgeFactory) CreateChuanchuan() ChuanchuanProduct {
	if e == nil {
		return nil
	}
	c := &YuLin{}
	return c
}
