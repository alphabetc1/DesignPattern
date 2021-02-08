package main

func main() {
	a := &Chicken{}
	a.s = a
	a.Talk()

	b := &Dog{}
	b.s = b
	b.Talk()
}
