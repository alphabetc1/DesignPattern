package main

import "fmt"

func main() {
	f := &Fund{StockOne{"Zhongxingguoji"},
		StockTwo{"Mayicaifu"},
		StockThree{"MaoTai"},
	}
	f.JiucaiBuy()
	fmt.Println()
	f.SmartBuy()
}
