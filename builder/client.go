package main

import "fmt"

func main() {
	thinPerson := PersonDirector{&PersonThin{"ShuFen"}}
	thinPerson.CreatePerson()

	fmt.Println()

	fatPerson := PersonDirector{&PersonFat{"LikeFlower"}}
	fatPerson.CreatePerson()
}
