package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func login() bool {
	var colors []string
	for i := 0; i < 10; i++ {
		colors = getlineColors(562, 724, 764)

		if compareColors(colors, "f5a322", "ffaa23") {
			robotgo.Move(642, 737)
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
		colors = getlineColors(602, 738, 609)
		if compareColors(colors, "ffffff", "c3c3c3") {
			robotgo.Move(669, 598)
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
		colors = getlineColors(1135, 1226, 680)
		if compareColor(colors, "037dd6") {
			robotgo.Move(1199, 680) // connect
			robotgo.Sleep(2)
			robotgo.Click("left")
			fmt.Println("click connect")
			robotgo.Sleep(8)
			return true
		}
		robotgo.Sleep(2)
	}
	return false
}
