package main

func main() {
	jingLi := &CommonManager{"张经理", nil}
	zongJian := &Majordomo{"洛总监", nil}
	zongJingLi := &GeneralManager{"马总经理", nil}
	jingLi.SetSuperior(zongJian)
	zongJian.SetSuperior(zongJingLi)

	request := &Request{"请假", "小王请假", 1}
	jingLi.RequestApplications(request)

	request2 := &Request{"请假", "小蔡请假", 5}
	jingLi.RequestApplications(request2)

	request3 := &Request{"加薪", "小王加薪", 500}
	jingLi.RequestApplications(request3)

	request4 := &Request{"加薪", "小李加薪", 5000}
	jingLi.RequestApplications(request4)
}
