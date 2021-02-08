package main

func main() {
	qiantaiMM := &ConerteSubject{}
	qiantaiMM.Attach(&ConcreteObserver{"小王", "摸鱼", qiantaiMM})
	qiantaiMM.Attach(&ConcreteObserver{"小鲲", "唱跳RAP", qiantaiMM})

	qiantaiMM.subjectState = "老板回来了，认真工作！"
	qiantaiMM.Notify()
}
