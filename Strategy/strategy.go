package main

import "fmt"

type CashSuper interface {
	acceptCash(float64)
}

type CashNormal struct {
}

func (c *CashNormal) acceptCash(money float64) {
	fmt.Println("原价:", money)
}

type CashRebate struct {
	rebate float64
}

func (c *CashRebate) acceptCash(money float64) {
	fmt.Println("打", c.rebate*10, "折:", int(money*c.rebate))
}

type CashReturn struct {
	needCondition float64
	needReturn    float64
}

func (c *CashReturn) acceptCash(money float64) {

	if money >= c.needCondition {
		money -= float64(int(money/c.needCondition)) * c.needReturn
	}

	fmt.Println("满", c.needCondition, "减", c.needReturn, ":", money)
}

type CashContext struct {
	CashSuper
}

func NewCashContext(str string) *CashContext {
	context := new(CashContext)
	switch str {
	case "满减":
		context.CashSuper = &CashReturn{300, 100}
	case "打折":
		context.CashSuper = &CashRebate{0.8}
	default:
		context.CashSuper = new(CashNormal)
	}
	return context
}
