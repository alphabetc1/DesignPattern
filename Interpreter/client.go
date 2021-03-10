package main

func main() {
	context := Context{"test"}
	list := []AbstractExpression{}

	list = append(list, new(TerminalExpression))
	list = append(list, new(TerminalExpression))
	list = append(list, new(NonterminalExpression))

	for _, val := range list {
		val.Interpret(&context)
	}
}
