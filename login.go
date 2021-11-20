package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func login() bool {
	fmt.Println("-------------start login")
	var colors []string
	for i := 0; i < 10; i++ {
		colors = getlineColors(570, 750, 740)

		if compareColors(colors, "f5a322", "ffaa23") {
			robotgo.Move(750, 740)
			robotgo.Click("left") // connect wallet
			fmt.Println("click connect wallet")
			robotgo.Sleep(2)
			continue
		}
		robotgo.Sleep(2)
		if i == 10 {
			return false
		}
	}
	for i := 0; i < 10; i++ {
		colors = getlineColors(672, 747, 588)
		if compareColors(colors, "ffffff", "c3c3c3") {
			robotgo.Move(747, 588)
			robotgo.Sleep(2)
			robotgo.Click("left") // metamask
			fmt.Println("click meta mask")
			robotgo.Sleep(8)
			continue
		}
		if i == 10 {
			return false
		}
	}

	for i := 0; i < 10; i++ {
		colors = getlineColors(1151, 1219, 647)
		if compareColor(colors, "037dd6") {
			robotgo.Move(1219, 647) // connect
			robotgo.Sleep(2)
			robotgo.Click("left")
			fmt.Println("click connect")
			robotgo.Sleep(8)
			return true
			fmt.Println("-------------start login")
		}
		robotgo.Sleep(2)
	}
	return false
}
