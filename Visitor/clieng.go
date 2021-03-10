package main

func main() {
	o := new(ObjectStructure)
	o.Attach(&Man{"男人"})
	o.Attach(&Woman{"女人"})

	o.Display(&Success{"成功"})
	o.Display(&Fail{"失败"})
}
