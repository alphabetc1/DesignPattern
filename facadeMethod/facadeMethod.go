/*
facade 外观模式
定义：为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得子系统易用。
作者：   alphabet
日期：   20201217
*/
package main

import "fmt"

type StockOne struct {
	name string
}

func (s *StockOne) Buy() {
	fmt.Println("Buy ", s.name)
}

type StockTwo struct {
	name string
}

func (s *StockTwo) Buy() {
	fmt.Println("Buy ", s.name)
}

type StockThree struct {
	name string
}

func (s *StockThree) Buy() {
	fmt.Println("Buy ", s.name)
}

type Fund struct {
	st1 StockOne
	st2 StockTwo
	st3 StockThree
}

func (f *Fund) JiucaiBuy() {
	fmt.Println("韭菜买法")
	f.st1.Buy()
	f.st2.Buy()
}

func (f *Fund) SmartBuy() {
	fmt.Println("聪明人买法")
	f.st3.Buy()
}
