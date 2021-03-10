package main

func main() {
	a := &Chicken{}
	a.Talk(a)

	b := &Dog{}
	b.Talk(b)
}
