package main

import "fmt"

func main() {
	root := NewConcreateCompany("北京总公司")
	root.Add(NewHRDepartment("总公司人力资源部"))
	root.Add(NewFinanceDepartment("总公司财务部"))

	compA := NewConcreateCompany("上海华东分公司")
	compA.Add(NewHRDepartment("上海华东分公司人力资源部"))
	compA.Add(NewFinanceDepartment("上海华东分公司财务部"))
	root.Add(compA)

	compB := NewConcreateCompany("南京办事处")
	compB.Add(NewHRDepartment("南京办事处人力资源部"))
	compB.Add(NewFinanceDepartment("南京办事处财务部"))
	compA.Add(compB)

	compC := NewConcreateCompany("杭州办事处")
	compC.Add(NewHRDepartment("杭州办事处人力资源部"))
	compC.Add(NewFinanceDepartment("杭州办事处财务部"))
	compA.Add(compC)

	fmt.Println("\n结构图：")
	root.Display(1)

	fmt.Println("\n职责：")
	root.LineOfDuty()
}
