package main

func main() {
	dota := Game{"Dota"}
	excel := Office{"Excel"}

	thinkPad := NewThinkPad(&dota, "Thinkpad")
	dell := NewDell(&excel, "Dell")
	thinkPad.Run()
	dell.Run()
}
