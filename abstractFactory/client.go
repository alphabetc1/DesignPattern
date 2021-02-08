package main

import "fmt"

func main() {
	var f AbstractFactory
	f = new(YigeFactory)

	h := f.CreateHuoguo()
	c := f.CreateChuanchuan()
	h.HuoguoInfo()
	c.ChuanchuanInfo()

	fmt.Println("========")

	f = new(ErgeFactory)

	h = f.CreateHuoguo()
	c = f.CreateChuanchuan()
	h.HuoguoInfo()
	c.ChuanchuanInfo()
}
