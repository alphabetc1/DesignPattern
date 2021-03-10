package main

import "fmt"

func main() {
	dinner := &Dinner{}
	first := Food{"第一道菜：上汤娃娃菜"}
	second := Food{"第二道菜：清炒西兰花"}
	third := Food{"第三道菜：虾仁滑蛋"}

	dinner.Add(first)
	dinner.Add(second)
	dinner.Add(third)

	concreteIt := dinner.CreateIterator()

	for iterator := concreteIt.First(); iterator != nil; iterator = concreteIt.Next() {
		fmt.Println(iterator)
	}
}
