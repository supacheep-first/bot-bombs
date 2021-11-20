package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func bot() bool {
	fmt.Println("----- start bot")

	robotgo.MouseSleep = 100
	robotgo.Sleep(5)
	robotgo.Move(667, 560) //treasure mode
	robotgo.Click("left")
	fmt.Println("click treasure mode")
	robotgo.Sleep(10)

	for i := 0; i < 15; i++ {

		ctl := robotgo.GetPixelColor(919, 616)
		ctr := robotgo.GetPixelColor(246, 376)
		cbt := robotgo.GetPixelColor(1032, 364)
		cbr := robotgo.GetPixelColor(1012, 700)

		if ctl == "150f1b" && ctr == "150f1b" && cbt == "150f1b" && cbr == "150f1b" {
			fmt.Errorf("found manual , unknow, something error set me hack to home page")
			return false
		}

		robotgo.Move(629, 733) // next map
		robotgo.Sleep(2)
		robotgo.Click("left")
		robotgo.Sleep(58) // 1min
		// robotgo.Sleep(6) // 6s

		if i%5 == 0 {
			robotgo.Move(215, 269) //back to home
			robotgo.Click("left")
			fmt.Println("click back to home")
			robotgo.Sleep(5)

			robotgo.Move(667, 560) //treasure mode
			robotgo.Click("left")
			fmt.Println("click treasure mode")
			robotgo.Sleep(10)
		}
	}

	// last click next map
	robotgo.Click("left")
	robotgo.Sleep(10)

	robotgo.Move(215, 269) //back to home
	robotgo.Click("left")
	fmt.Println("click back to home")
	robotgo.Sleep(5)

	fmt.Println("----- end bot")
	return true
}
