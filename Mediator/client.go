package main

func main() {
	UNSC := NewUnitedNationsCouncil("联合国")

	usa := &USA{"美国", UNSC}
	china := &China{"中国", UNSC}
	europe := &Europe{"欧洲", UNSC}

	UNSC.AddCountry(usa)
	UNSC.AddCountry(china)
	UNSC.AddCountry(europe)

	usa.Declare("新冠病毒是起源于中国的！")
	china.Declare("众志成城合作抗议")
	europe.Declare("求疫苗")
}
