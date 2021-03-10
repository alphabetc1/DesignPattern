package main

func main() {
	jiaojiao := SchoolGirl{"娇娇"}

	jiayi := Pursuit{"卓假意", jiaojiao}
	daili := Proxy{"戴笠", jiayi}

	daili.GiveDolls()
	daili.GiveFlowers()
	daili.GiveChocolate()
}
