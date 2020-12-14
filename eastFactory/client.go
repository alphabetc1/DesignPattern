package main

import "fmt"

func main() {
	var operationFactory OperationFactory
	oper := operationFactory.createOperate("+")
	oper.SetNumberA(1)
	oper.SetNumberB(2)

	result := oper.GetResult()
	fmt.Println(result)
}
