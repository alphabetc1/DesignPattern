package main

func main() {
	xiaoWang := Barbecuer{"小王"}
	xiaoLi := Barbecuer{"小李"}
	bakeMuttonCommand1 := BakeMuttonCommand{xiaoWang}
	bakeMuttonCommand2 := BakeMuttonCommand{xiaoWang}
	bakeChickenWingCommand1 := BakeChickenWingCommand{xiaoLi}
	girl := &Waiter{}

	girl.SetOrder(&bakeMuttonCommand1)
	girl.SetOrder(&bakeMuttonCommand2)
	girl.SetOrder(&bakeChickenWingCommand1)

	girl.Notify()
}
