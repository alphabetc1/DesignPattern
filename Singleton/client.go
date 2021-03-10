package main

import "fmt"

func main() {
	sin1 := GetSingleton()
	sin2 := GetSingleton()

	if sin1 != sin2 {
		fmt.Println("Wonderful code")
	}
}
