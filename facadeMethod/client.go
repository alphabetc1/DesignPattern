package main

func main() {
	f := &Fund{StockOne{"Zhongxingguoji"},
		StockTwo{"Mayicaifu"},
		StockThree{"MaoTai"},
	}
	f.JiucaiBuy()
	f.SmartBuy()
}
