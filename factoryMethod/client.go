package main

import "fmt"

func main() {
	operFactory := AddFactory{}
	var oper OperationFunc = operFactory.CreateOperation()
	oper.SetNumberA(1)
	oper.SetNumberB(2)

	result := oper.GetResult()
	fmt.Println(result)
}
