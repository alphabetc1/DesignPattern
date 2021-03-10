package main

import "fmt"

type Barbecuer struct {
	name string
}

type Command interface {
	ExcuteCommand()
}

type BakeMuttonCommand struct {
	barbecuer Barbecuer
}

func (b *BakeMuttonCommand) Name() string {
	return b.barbecuer.name
}

func (b *BakeMuttonCommand) ExcuteCommand() {
	fmt.Println("厨师", b.Name(), "在烤肉")
}

type BakeChickenWingCommand struct {
	barbecuer Barbecuer
}

func (b *BakeChickenWingCommand) Name() string {
	return b.barbecuer.name
}

func (b *BakeChickenWingCommand) ExcuteCommand() {
	fmt.Println("厨师", b.Name(), "在烤鸡翅")
}

type Waiter struct {
	orders []Command
}

func (w *Waiter) SetOrder(command Command) {
	if w == nil {
		return
	}
	w.orders = append(w.orders, command)
}

func (w *Waiter) CancelOrder(command Command) {
	if w == nil {
		return
	}
	for i, order := range w.orders {
		if order == command {
			w.orders = append(w.orders[:i], w.orders[i+1:]...)
			return
		}
	}
}

func (w *Waiter) Notify() {
	if w == nil {
		return
	}
	for _, order := range w.orders {
		order.ExcuteCommand()
	}
}
