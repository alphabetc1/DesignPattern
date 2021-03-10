package main

import "fmt"

func main() {
	xw := &Person{"xiaoWang"}
	xw.Show()

	fmt.Println("=======第一种装扮=======")
	trouser := new(BigTrouser)
	tshirt := new(Tshirt)

	trouser.SetDecorator(xw)
	tshirt.SetDecorator(trouser)
	tshirt.Show()

	fmt.Println("=======第二种装扮=======")
	miniskirt := new(Miniskirt)
	suit := new(Suit)

	miniskirt.SetDecorator(xw)
	suit.SetDecorator(miniskirt)
	suit.Show()
}
