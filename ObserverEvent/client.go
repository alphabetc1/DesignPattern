package main

func main() {
	boss := &Boss{"我胡汉三回来了！", []*update{}}
	xiaoWang := &StockObserver{"小王"}
	xiaoKun := &NBAObserver{"小鲲"}

	var StockEvent update = xiaoWang.CloseStockMarket
	var NBAEvent update = xiaoKun.CloseNBADirectSeeding

	boss.AddEventHandler(&StockEvent)
	boss.AddEventHandler(&NBAEvent)

	boss.Notify()
}
