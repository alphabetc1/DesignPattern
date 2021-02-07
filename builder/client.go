package main

import "fmt"

func main() {
	thin := PersonThin{"ShuFen"}
	fat := PersonFat{"LikeFlower"}

	directorThin := PersonDirector{&thin}
	directorThin.CreatePerson()

	fmt.Println()

	directorFat := PersonDirector{&fat}
	directorFat.CreatePerson()
}
