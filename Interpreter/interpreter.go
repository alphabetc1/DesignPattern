package main

import "fmt"

type Context struct {
	text string
}

type AbstractExpression interface {
	Interpret(*Context)
}

type TerminalExpression struct {
}

func (t *TerminalExpression) Interpret(context *Context) {
	if t == nil {
		return
	}
	context.text = context.text[:len(context.text)-1]
	fmt.Println("终结符解释器：", context.text)
}

type NonterminalExpression struct {
}

func (n *NonterminalExpression) Interpret(context *Context) {
	if n == nil {
		return
	}
	fmt.Println("非终结符解释器：", context.text)
}
