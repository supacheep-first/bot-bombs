package main

import "github.com/go-vgo/robotgo"

func main() {
	// setup infinity loop
	// init time
	robotgo.Sleep(5)

	// idk why set but im don't forget set it!!!!
	robotgo.MouseSleep = 100

	infinity := 1
	for {
		// re tab every 5 hr.
		if infinity%15 == 0 {
			doRefresh()
		}
		if infinity%15 == 0 || infinity == 1 {
			if !login() {
				doRefresh()
				continue
			}
		}

		if !selectHero2() {
			infinity = 0
			continue
		}
		if !bot() {
			infinity = 0
			continue
		}

		infinity++
	}

}

// func reOpenGame() {
// 	closeTab()
// 	openTabGame()
// }
