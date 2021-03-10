package main

func main() {
	liXiaoYao := &Role{"李逍遥", 100, 100, 100}
	liXiaoYao.StateDisplay()

	stateManager := &RoleManager{liXiaoYao.SaveState()}

	liXiaoYao.Fight()
	liXiaoYao.StateDisplay()

	liXiaoYao.RecoveryState(stateManager.state)
	liXiaoYao.StateDisplay()
}
