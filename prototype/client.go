package main

func main() {
	a := NewResume()
	a.SetPersonalInfo("小王", "男", "22")
	a.SetWorkExperience("1998-2000", "apple")
	a.Display()

	b := a.Clone()
	b.SetPersonalInfo("小李", "男", "35")
	b.SetWorkExperience("2000-2014", "oldNationalBusiness")
	b.Display()
	a.Display()

	c := b.CloneWithNewExperience("2014-2018", "banana")
	c.SetPersonalInfo("小张", "女", "26")
	c.Display()
	b.Display()
}
