/*
Memento 备忘录模式
定义：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
适用场景：功能比较复杂，但需要维护和记录属性历史的类，也可用于撤销。
作者：   alphabetc1
日期：   20210219
*/
package main

import "fmt"

type Role struct {
	name          string
	vit, atk, def int
}

func (r *Role) Fight() {
	r.vit = 0
}

func (r *Role) StateDisplay() {
	fmt.Println("角色", r.name, "攻击力：", r.atk, "防御力：", r.def, "生命力：", r.vit)
}

func (r *Role) SaveState() *RoleMemento {
	return &RoleMemento{r.vit, r.atk, r.def}
}

func (r *Role) RecoveryState(m *RoleMemento) {
	r.vit, r.atk, r.def = m.vit, m.atk, m.def
}

type RoleMemento struct {
	vit, atk, def int
}

type RoleManager struct {
	state *RoleMemento
}
