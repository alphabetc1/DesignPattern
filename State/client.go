package main

func main() {
	w := NewWork()
	w.WriteProgram()
	w.SetHour(12)
	w.WriteProgram()
	w.SetHour(19)
	w.WriteProgram()
	w.SetHour(21)
	w.WriteProgram()
}
