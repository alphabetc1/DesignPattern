package main

func main() {
	context := NewCashContext("正常")
	context.acceptCash(998)

	context = NewCashContext("满减")
	context.acceptCash(998)

	context = NewCashContext("打折")
	context.acceptCash(998)
}
