package main

type OperationFunc interface {
	SetNumberA(float32)
	SetNumberB(float32)
	GetResult() float32
}

type Operation struct {
	NumberA float32
	NumberB float32
}

func (o *Operation) SetNumberA(num float32) {
	o.NumberA = num
}

func (o *Operation) SetNumberB(num float32) {
	o.NumberB = num
}

type OperationAdd struct {
	Operation
}

func (o OperationAdd) GetResult() float32 {
	return o.NumberA + o.NumberB
}

type OperationSub struct {
	Operation
}

func (o OperationSub) GetResult() float32 {
	return o.NumberA - o.NumberB
}

type OperationFactory struct {
}

func (o OperationFactory) createOperate(str string) OperationFunc {
	var oper OperationFunc
	switch str {
	case "+":
		oper = &OperationAdd{}
		break
	case "-":
		oper = &OperationSub{}
		break
	}

	return oper
}
